<?php

use Spiral\Goridge;
use Spiral\Goridge\RelayInterface as Relay;

require "vendor/autoload.php";

$rpc = new Goridge\RPC(new Goridge\SocketRelay("127.0.0.1", 6001));
//读取文件
$file = file_get_contents($_FILES["file"]["tmp_name"]);
//直接发送二进制文件
$data = $rpc->call("Excel.Decode", $file, Relay::PAYLOAD_RAW);

print_r($data);