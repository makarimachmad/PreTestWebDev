-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 11 Okt 2020 pada 03.56
-- Versi server: 10.1.37-MariaDB
-- Versi PHP: 7.3.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `video`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `kamutube`
--

CREATE TABLE `kamutube` (
  `id` int(11) NOT NULL,
  `src` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `artist` varchar(255) NOT NULL,
  `sifat` tinyint(1) NOT NULL,
  `deskripsi` varchar(600) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `kamutube`
--

INSERT INTO `kamutube` (`id`, `src`, `title`, `artist`, `sifat`, `deskripsi`) VALUES
(1, '@/assets/videos/Shanna Shannon - Kamu dan Kenangan  OST Habibie Ainun 3  (Cover).webm', 'Kamu dan Kenangan  OST Habibie Ainun 3  (Cover)', 'Shanna Shannon', 1, 'Buat aku lagu ini bukan hanya sekedar lagu cinta orang dewasa melainkan lagu yang bisa dinikmati oleh semua usia karena mempunyai banyak arti. \r\n\r\nAku bersyukur karena lewat lagu ini aku bisa belajar banyak hal. Aku belajar mengucap syukur untuk orang-orang yang kita sayangi dan terus berbagi selama masih ada kesempatan, belajar tentang kesetiaan, tentang kasih sayang, tentang ketulusan, dan belajar untuk selalu bersyukur dan menghargai kebersamaan.  \r\n\r\nSelain itu, aku pengagum Pak Habibie dan Ibu Ainun dengan kisahnya yang memberi banyak teladan buat banyak orang.\r\n\r\nSemoga cover saya ini bi'),
(2, '@/assets/videos/Hindia - Secukupnya (Lyric Video) - OST. Nanti Kita Cerita Tentang Hari Ini.webm', 'Secukupnya (Lyric Video) - OST. Nanti Kita Cerita Tentang Hari Ini', 'Hindi', 1, 'Listen and stream on: https://backl.ink/secukupnyahindia.oyd\r\n\r\nSong Title:  Secukupnya\r\nSinger:  Hindia\r\nMusic by Adhe Arrio\r\nLyrics by Baskara Putra\r\nProduced & Arranged by Adhe Arrio\r\nVocals Directed & Arranged by Baskara Putra \r\nMixed by Wisnu Ikhsantama W\r\nMastered by Wisnu Ikhsantama W\r\nExecutive Producer/Records Company : Sun Eater\r\n\r\nLyric Video by\r\nPhotographer: Jozz Felix & Andreansyah Dimas W. P. G.\r\nAnimator: Arief Khoirul Alim'),
(3, '@/assets/videos/Pamungkas - Kenangan Manis (Lyrics Video).webm', 'Kenangan Manis (Lyrics Video)', 'Pamungkas', 1, 'Produced, Written, Recorded, Mixed and Mastered by Pamungkas.\r\n\r\nWalk The Talk is available on Digital Platforms.\r\nFollow Pamungkas on Instagram; \r\nhttp://www.instagram.com/pamunqkas'),
(4, '@/assets/videos/Fourtwnty - Nematomorpha (Official Music Video).webm', 'Nematomorpha (Official Music Video)', 'Fourtwenty', 1, 'A Visual by : Seadaadanya\r\n\r\nCast  : \r\nAri Lesmana\r\nNuwi\r\nPrimanda Ridho\r\nAndiarmand\r\nRyan Maulana\r\n\r\nDirected by : Vicko Jonanda\r\nDirector of Photography/ ArtworkDesign : Reza Vebrian\r\nEditor : Liemena ( Pau )\r\nMake Up Artist : Adeline Chandra\r\nWardrobe Stylist : \r\nJade Daliz\r\nAdeline Chandra\r\nBTS Videographer : \r\nNurul Fajar\r\nIkhwanul Muslimin\r\n\r\nThanks to :\r\nRoby Satria\r\nEdo Gordo\r\nAdriansyah Bustami\r\nUkien\r\nDedi Ob\r\nKaryaintijiwa\r\nZinanyamaniku\r\nAll crew');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `kamutube`
--
ALTER TABLE `kamutube`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `kamutube`
--
ALTER TABLE `kamutube`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
