<?php

use Spiral\Goridge;

require "vendor/autoload.php";

$rpc = new Goridge\RPC(new Goridge\SocketRelay("127.0.0.1", 6001));

$data = [];


for ($i = 0; $i < 6; $i++) {
    $data[] = ['yaecho', 'https://yaecho.net', '2019-01-01', 'haha', '12:00', '2131245666', 2019];
}


$file = $rpc->call("Excel.Encode", $data);

header('Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet');
header('Content-Disposition: attachment; filename=test.xlsx');
header('Expires: 0');
header('Cache-Control: must-revalidate');
header('Pragma: public');
header('Content-Length: ' . strlen($file));

ob_clean();
flush();
echo $file;
exit;