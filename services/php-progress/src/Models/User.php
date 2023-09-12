<?php

namespace Kubecamp\Progress\Models;

use DateTime;
use Exception;

use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\ODM\MongoDB\Mapping\Annotations as ODM;
use JsonSerializable;
use Kubecamp\Progress\Database\MongoDocumentManager as DocumentManager;
use ReturnTypeWillChange;

/** @ODM\Document*/
class User implements JsonSerializable
{
  /** @ODM\Id */
  private $id;

  /** @ODM\Field(type="string") */
  private $userId;

  /** @ODM\EmbedMany(targetDocument=Progress::class) */
  private $progress;

  public function __construct()
  {
    $this->progress = new ArrayCollection();
  }

  #[ReturnTypeWillChange]
  public function jsonSerialize()
  {
    return [
      'id' => $this->id,
      'userId' => $this->userId,
      'progress' => $this->progress->map(function (Progress $progress) {
        return $progress->jsonSerialize();
      })->toArray()
    ];
  }

  public static function getProgressByUserId($userId)
  {
    $dm = DocumentManager::getDocumentManager();
    $user = $dm->getRepository(User::class)->findOneBy(['userId' => $userId]);
    return $user;
  }

  public static function appendProgress($userId, $lessonId)
  {
    $dm = DocumentManager::getDocumentManager();
    error_log("⛑️ Appending progress for user $userId on lesson $lessonId");

    // Find or create the user
    $user = $dm->getRepository(User::class)->findOneBy(['userId' => $userId]);

    if (!$user) {
      $user = new User();
      $user->userId = $userId;
      $dm->persist($user);
    }

    // Add the new progress entry
    $progress = new Progress(new DateTime(), $lessonId);
    $user->addProgress($progress);

    // Flush changes to database
    $dm->flush();
  }

  public function addProgress(Progress $progress)
  {
    $existingEntries = $this->progress->filter(function (Progress $existingProgress) use ($progress) {
      return $existingProgress->getLessonId() == $progress->getLessonId();
    });

    if ($existingEntries->count() < 1) {
      $this->progress->add($progress);
    } else {
      throw new Exception("Cannot add more than two progress entries with the same lessonId");
    }
  }
}

/** @ODM\EmbeddedDocument */
class Progress implements JsonSerializable
{
  /** @ODM\Field(type="date") */
  private $completionDate;

  /** @ODM\Field(type="int") */
  private $lessonId;

  public function __construct($date, $lessonId)
  {
    $this->completionDate = $date;
    $this->lessonId = $lessonId;
  }

  #[ReturnTypeWillChange]
  public function jsonSerialize()
  {
    return [
      'completionDate' => $this->completionDate->format('Y-m-d H:i:s'),  // Format date to a string representation
      'lessonId' => $this->lessonId
    ];
  }


  public function getLessonId()
  {
    return $this->lessonId;
  }

  public function getCompletionDate()
  {
    return $this->completionDate;
  }
}
