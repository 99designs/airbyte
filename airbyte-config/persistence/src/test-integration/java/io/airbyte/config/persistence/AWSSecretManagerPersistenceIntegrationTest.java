/*
 * Copyright (c) 2021 Airbyte, Inc., all rights reserved.
 */

package io.airbyte.config.persistence;

import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertEquals;

import io.airbyte.config.persistence.split_secrets.AWSSecretManagerPersistence;
import io.airbyte.config.persistence.split_secrets.SecretCoordinate;
import java.io.IOException;
import org.apache.commons.lang3.RandomUtils;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import java.util.List;

/**
 * Triggered as part of integration tests in CI. It uses credentials in Github to connect to the
 * integration testing AWS project.
 */
public class AWSSecretManagerPersistenceIntegrationTest {

  private AWSSecretManagerPersistence persistence;
  private String coordinateBase;

  @BeforeEach
  void setUp() {
    persistence = AWSSecretManagerPersistence.getEphemeral();
    coordinateBase = "AWSSecretManagerPersistenceIntegrationTest_coordinate_" + RandomUtils.nextInt() % 20000;
  }

  @AfterEach
  void tearDown() throws IOException {
    // TODO: Delete all secrets that share the coordinateBase
    List<SecretCoordinate> secrets = persistence.listSecretVersions(coordinateBase);
    System.out.println(secrets);

    secrets.forEach(secret -> persistence.delete(secret));
  }

  @Test
  void testReadWriteUpdate() {
    final var coordinate1 = new SecretCoordinate(coordinateBase, 1);

    // try reading non-existent value
    final var firstRead = persistence.read(coordinate1);
    assertTrue(firstRead.isEmpty());

    // write
    final var firstPayload = "abc";
    persistence.write(coordinate1, firstPayload);
    final var secondRead = persistence.read(coordinate1);
    assertTrue(secondRead.isPresent());
    assertEquals(firstPayload, secondRead.get());
  }

}
