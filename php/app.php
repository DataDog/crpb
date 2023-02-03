<?php

$port = getenv('PORT');

if ($port === false) {
    $port = 8080;
}

$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);

socket_bind($socket, 0, $port);
socket_listen($socket);

while(true) {
    $client = socket_accept($socket);
    $request = '';
    socket_read($client, 1024, PHP_NORMAL_READ);

    $response = "HTTP/1.1 200 OK\r\n";
    $response .= "Content-Type: text/plain\r\n";
    $response .= "Content-Length: 12\r\n";
    $response .= "\r\n";
    $response .= "Hello World!";

    socket_write($client, $response);
    socket_close($client);
}

socket_close($socket);