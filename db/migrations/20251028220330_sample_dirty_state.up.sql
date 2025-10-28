-- ==============================================
-- Migration: sample_dirty_state (UP)
-- Database : belajar_golang_database_migration
-- Created  : Tue Oct 28 10:03:30 PM WITA 2025
-- ==============================================

-- Tulis SQL untuk UP migration di bawah ini:
CREATE TABLE correct
(
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
)ENGINE = InnoDB;

CREATE TABLE wrong
(
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
)ENGINE = InnoDB;
