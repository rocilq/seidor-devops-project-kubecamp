<?php

namespace Kubecamp\Progress\Controllers;

use Exception;
use Kubecamp\Progress\Models\User;
use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class ProgressController
{

  public static function getProgress(Request $request, Response $response, $name)
  {
    $userId = $request->getAttribute('id');

    try {
      error_log("⛑️ Getting progress for user $userId");
      $user = User::getProgressByUserId($userId);
      error_log(json_encode($user));
      $response->getBody()->write(json_encode($user));
      return $response->withHeader('Content-Type', 'application/json');
    } catch (Exception $e) {
      $response->getBody()->write(json_encode(['error' => "Not found"]));
      return $response->withStatus(404);
    }
  }

  public static function appendProgress(Request $request, Response $response, $name)
  {
    $userId = $request->getAttribute('id');
    $body = $request->getBody();
    $data = json_decode($body, true);
    $lessonId = $data['lessonId'] ?? null;

    if (!$userId || !$lessonId || !is_numeric($lessonId)) {
      $response->getBody()->write(json_encode(['error' => 'Invalid request']));
      return $response->withStatus(400)->withHeader('Content-Type', 'application/json');
    }

    try {
      User::appendProgress($userId, $lessonId);
      $user = User::getProgressByUserId($userId);
      $response->getBody()->write(json_encode(['user' => $user]));
      return $response->withStatus(201)->withHeader('Content-Type', 'application/json');
    } catch (Exception $e) {
      error_log($e);
      $response->getBody()->write(json_encode(['error' => $e->getMessage()]));
      return $response->withStatus(500)->withHeader('Content-Type', 'application/json');
    }
  }
}
