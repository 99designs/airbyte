/*
 * Copyright (c) 2021 Airbyte, Inc., all rights reserved.
 */

package io.airbyte.config.persistence.split_secrets;

import io.airbyte.commons.lang.Exceptions;
import java.util.List;
import java.util.Optional;
import java.util.function.Supplier;
import java.util.function.Supplier;
import java.util.stream.Collectors;
import software.amazon.awssdk.services.secretsmanager.SecretsManagerClient;
import software.amazon.awssdk.services.secretsmanager.model.CreateSecretRequest;
import software.amazon.awssdk.services.secretsmanager.model.DeleteSecretRequest;
import software.amazon.awssdk.services.secretsmanager.model.Filter;
import software.amazon.awssdk.services.secretsmanager.model.GetSecretValueRequest;
import software.amazon.awssdk.services.secretsmanager.model.GetSecretValueResponse;
import software.amazon.awssdk.services.secretsmanager.model.ListSecretsRequest;
import software.amazon.awssdk.services.secretsmanager.model.ListSecretsResponse;
import software.amazon.awssdk.services.secretsmanager.model.PutResourcePolicyRequest;
import software.amazon.awssdk.services.secretsmanager.model.ResourceNotFoundException;
import software.amazon.awssdk.services.sts.StsClient;

public class AWSSecretManagerPersistence implements SecretPersistence {

  private final Supplier<SecretsManagerClient> secretManagerClientSupplier;
  private final Supplier<StsClient> stsClientSupplier;

  /**
   * A private constructor class which will be used by getEphemeral and
   * getLongLived to create a new instance of the secret persistence.
   */
  private AWSSecretManagerPersistence() {
    // The SecretsManagerClient will pull in credentials and configuration from
    // the Default Credential Provider Chain so we don't actually need to bring
    // in the any configuration.
    this.secretManagerClientSupplier = () -> getSecretManagerClient();
    this.stsClientSupplier = () -> getStsClient();
  }

  public static SecretsManagerClient getSecretManagerClient() {
    return SecretsManagerClient.builder().build();
  }

  public static StsClient getStsClient() {
    return StsClient.builder().build();
  }

  /**
   * Creates a persistence with a relatively short TTL for stored secrets. Used
   * for temporary operations such as check/discover operations where we need to
   * use secret storage to communicate from the server to Temporal, but where we
   * don't want to maintain the secrets indefinitely.
   */
  public static AWSSecretManagerPersistence getEphemeral() {
    return new AWSSecretManagerPersistence();
  }

  /**
   * Creates a persistence with an infinite TTL for stored secrets. Used for
   * source/destination config secret storage.
   */
  public static AWSSecretManagerPersistence getLongLived() {
    return new AWSSecretManagerPersistence();
  }

  @Override
  public Optional<String> read(final SecretCoordinate coordinate) {
    try (final var client = secretManagerClientSupplier.get()) {
      String secretName = coordinate.getFullCoordinate();

      GetSecretValueRequest request = GetSecretValueRequest
        .builder()
        .secretId(secretName)
        .build();

      GetSecretValueResponse response = client
        .getSecretValue(request);

      return Optional.of(response.secretString());
    } catch (final ResourceNotFoundException e) {
      return Optional.empty();
    }
  }

  @Override
  public void write(final SecretCoordinate coordinate, final String payload) {
    try (final var client = secretManagerClientSupplier.get()) {
      String secretName = coordinate.getFullCoordinate();

      CreateSecretRequest request = CreateSecretRequest.builder()
        .name(secretName)
        .description("This secret was created by the Airbyte.")
        .secretString(payload)
        .build();

      client.createSecret(request);

      this.allowDeletion(secretName);
    }
  }

  /**
   * List all the versions of a particular SecretCoordinate.
   */
  public List<SecretCoordinate> list(final String coordinateBase) {
    try (final var client = secretManagerClientSupplier.get()) {
      Filter filter = Filter
        .builder()
        .key("name")
        .values(coordinateBase)
        .build();

      ListSecretsRequest request = ListSecretsRequest
        .builder()
        .filters(filter)
        .build();

      ListSecretsResponse response = client.listSecrets(request);

      return response
        .secretList()
        .stream()
        .map(secret -> SecretCoordinate.fromFullCoordinate(secret.name()))
        .collect(Collectors.toList());
    }
  }

  /**
   * You can't just delete a secret in AWS Secret Manager, you first need to add
   * a resource policy which allows you to delete the secret.
   */
  private void allowDeletion(final String secretName) {
    try (final var secretManagerClient = secretManagerClientSupplier.get();
         final var stsClient = stsClientSupplier.get()) {
      String account = stsClient.getCallerIdentity().account();

      String policy = String.format(
        """
        {
          "Version":"2012-10-17",
          "Statement": {
            "Effect": "Allow",
            "Principal": {
              "AWS": "%s"
            },
            "Action": "secretsmanager:DeleteSecret",
            "Resource": "*"
          }
        }
        """,
        account
      );

      PutResourcePolicyRequest request = PutResourcePolicyRequest
        .builder()
        .secretId(secretName)
        .resourcePolicy(policy)
        .build();

      secretManagerClient.putResourcePolicy(request);
    }
  }

  /**
   * Delete a secret.
   */
  public void delete(final SecretCoordinate coordinate) {
    try (final var client = secretManagerClientSupplier.get()) {
      String secretName = coordinate.getFullCoordinate();

      DeleteSecretRequest request = DeleteSecretRequest
        .builder()
        .recoveryWindowInDays(Long.valueOf(7))
        .build();

      client.deleteSecret(request);
    }
  }

}
