<?php

use Doctrine\ODM\MongoDB\Repository\DocumentRepository;
use PHPUnit\Framework\TestCase;
use Kubecamp\Progress\Models\User;
use Kubecamp\Progress\Models\Progress;
use Kubecamp\Progress\Database\MongoDocumentManager as DocumentManager;

class UserTest extends TestCase
{
  public function testJsonSerialize()
  {
    $user = new User();
    $progress = new Progress(new DateTime(), 1);

    $expected = [
      'id' => null, // because it's not set in this test
      'userId' => null,
      'progress' => []
    ];

    $this->assertEquals($expected, $user->jsonSerialize());

    $user->addProgress($progress);

    $expected['progress'] = [$progress->jsonSerialize()];
    $this->assertEquals($expected, $user->jsonSerialize());
  }

  public function testAddProgressWithExistingLessonId()
  {
    $user = new User();
    $progress1 = new Progress(new DateTime(), 1);
    $progress2 = new Progress(new DateTime(), 1);

    $user->addProgress($progress1);
    $this->expectException(\Exception::class);
    $this->expectExceptionMessage('Cannot add more than two progress entries with the same lessonId');

    $user->addProgress($progress2);
  }
}

class ProgressTest extends TestCase
{
  public function testJsonSerialize()
  {
    $date = new DateTime();
    $lessonId = 1;

    $progress = new Progress($date, $lessonId);
    $expected = [
      'completionDate' => $date->format('Y-m-d H:i:s'),
      'lessonId' => $lessonId
    ];

    $this->assertEquals($expected, $progress->jsonSerialize());
  }
}
