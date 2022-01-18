-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 18 Jan 2022 pada 12.22
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.4.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test_sagara`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `product`
--

CREATE TABLE `product` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `price` int(8) NOT NULL,
  `image` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product`
--

INSERT INTO `product` (`id`, `name`, `description`, `price`, `image`, `created_at`, `updated_at`) VALUES
(11, 'Baju Keren Pria', 'Baju Keren Pria Produk Asli Indonesia', 80000, '/api/products/image/2883247955.jpg', '2022-01-18 10:02:40', '2022-01-18 10:05:10');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `fullname` varchar(255) NOT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `user`
--

INSERT INTO `user` (`id`, `fullname`, `username`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Eriawan Hidayat', 'eriawan', '$2a$12$2ClaMdfjAShZaaLSl/raj.E7r0PRZv4es4Dosg2mWiCwPlMAhgkFS', '2022-01-18 05:56:22', '2022-01-18 05:56:22'),
(2, 'Jane Doe', 'jane', '$2a$12$Rj2UbVgWCSp2ciqP5bqc1usGdiOx7Kp6Jv5xA7u7C9Gf0z3saODL2', '2022-01-18 05:59:42', '2022-01-18 05:59:42'),
(3, 'John Doe', 'john', '$2a$12$MPMUXhzC508yAfY1ZKPyrOT0cls0TqA5BblD8FrcdNY79ft0AlCjW', '2022-01-18 06:43:26', '2022-01-18 06:43:26'),
(4, 'Joni Doe', 'joni', '$2a$12$Xz2VLKrXFALAG.4j9nM4iewkzKUl6koOtzR9/lESHXPlvJOYGSbIu', '2022-01-18 06:45:10', '2022-01-18 06:45:10'),
(5, 'Jojon Doe', 'jojon', '$2a$12$DbkWQOudL5JpeaoNjDVN6e/8y6b4G4tvzIYtEAif9eeb2WAgzt0t2', '2022-01-18 06:46:49', '2022-01-18 06:46:49'),
(6, 'Jajat Doe', 'jajat', '$2a$12$myDUeE2D2FiVJpHLiRF5Uu1N6a0AgOvCwIE.ap6ZiCB9TNKQkbMxa', '2022-01-18 06:51:20', '2022-01-18 06:51:20'),
(7, 'Jeni Doe', 'jeni', '$2a$12$hq.ETCXSdb4HLGbUPerkWO7INAlVoJT9VUN5dMhvIHRrg73hcNGMC', '2022-01-18 06:51:57', '2022-01-18 06:51:57'),
(8, 'jijin Doe', 'jijin', '$2a$12$a0uConktdN2lgdENiI4iXeh9fD2dqkbxTM8qCucNBvUGpv0zIg.3i', '2022-01-18 06:52:15', '2022-01-18 06:52:15'),
(9, 'Dummy User', 'dummyuser1', '$2a$12$rQ7BrEkEmuDcwgW4zyG/u.xqGd58fvhTyrV7dkamGOdkIEf/xgLH6', '2022-01-18 09:58:09', '2022-01-18 09:58:09');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `product`
--
ALTER TABLE `product`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT untuk tabel `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
