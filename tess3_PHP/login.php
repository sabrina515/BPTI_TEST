<?php
// Konfigurasi database
$dbHost     = '127.0.0.1';
$dbUsername = 'root';  // Ganti dengan username MySQL Anda
$dbPassword = '';  // Ganti dengan password MySQL Anda
$dbName     = 'tes2';      // Ganti dengan nama database Anda

// Membuat koneksi ke database
$db = new mysqli($dbHost, $dbUsername, $dbPassword, $dbName);

// Cek koneksi
if ($db->connect_error) {
    die("Connection failed: " . $db->connect_error);
}

// Menangani form login
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $username = $db->real_escape_string($_POST['username']);
    $password = $db->real_escape_string($_POST['password']);

    // Mencari user di database
    echo password_hash($password, PASSWORD_BCRYPT, []);
    $query = "SELECT * FROM author WHERE username = '$username' AND password = '".password_hash($password, PASSWORD_BCRYPT,['cost' => 11])."'";
    $result = $db->query($query);

    if ($result->num_rows > 0) {
        echo "Login berhasil!";
        // Set session atau redirect user ke halaman lain
    } else {
        echo "Username atau password salah!";
    }
    $db->close();
}
?>