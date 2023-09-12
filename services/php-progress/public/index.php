<?php

require __DIR__ . '/../vendor/autoload.php';

use Kubecamp\Progress\Controllers\ProgressController;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;
use Slim\Factory\AppFactory;


$app = AppFactory::create();

$app->get('/', function (Request $request, Response $response, $args) {
  $response->getBody()->write("Hello world!");
  return $response;
});

$app->get('/progress/{id}', ProgressController::class . ':getProgress');
$app->post('/progress/{id}', ProgressController::class . ':appendProgress');

$app->run();
