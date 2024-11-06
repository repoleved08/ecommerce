-- Create Products Table
CREATE TABLE `products` (
    `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `image` VARCHAR(255) NOT NULL,
    `category` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `rating` INT NOT NULL,
    `num_reviews` INT NOT NULL DEFAULT 0,
    `price` DECIMAL(18, 2) NOT NULL,
    `count_in_stock` INT NOT NULL,
    `created_at` DATETIME DEFAULT (now()),
    `updated_at` DATETIME
);

-- Create Orders Table
CREATE TABLE `orders` (
    `id` INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `payment_method` VARCHAR(255) NOT NULL,
    `tax_price` DECIMAL(18, 2) NOT NULL,
    `shipping_price` DECIMAL(10, 2) NOT NULL,
    `total_price` DECIMAL(10, 2) NOT NULL,
    `created_at` DATETIME DEFAULT (now()),
    `updated_at` DATETIME
);

-- Create Order Items Table
CREATE TABLE `order_items` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `order_id` INT NOT NULL,
    `product_id` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `quantity` INT NOT NULL,
    `image` VARCHAR(255) NOT NULL,
    `price` DECIMAL(18, 2) NOT NULL
);

ALTER TABLE `order_items` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);
ALTER TABLE `order_items` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
