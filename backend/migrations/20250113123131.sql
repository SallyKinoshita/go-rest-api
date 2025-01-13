-- Create "sample_user" table
CREATE TABLE `sample_user` (
  `id` char(36) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `first_name_kana` varchar(255) NOT NULL,
  `last_name_kana` varchar(255) NOT NULL,
  `email_address` varchar(255) NOT NULL,
  `tel` varchar(255) NOT NULL,
  `birth_date` datetime NOT NULL,
  `postal_code` varchar(255) NOT NULL,
  `prefecture` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `street_and_number` varchar(255) NOT NULL,
  `building_and_room` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
