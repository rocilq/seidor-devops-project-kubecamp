<?

namespace Kubecamp\Progress\Database;

use Doctrine\Common\Annotations\AnnotationReader;
use Doctrine\ODM\MongoDB\DocumentManager;
use Doctrine\ODM\MongoDB\Configuration;
use Doctrine\ODM\MongoDB\Mapping\Driver\AnnotationDriver;
use MongoDB\Client;

class MongoDocumentManager
{
  private static $documentManager;

  static public function getDocumentManager()
  {
    if (self::$documentManager) {
      return self::$documentManager;
    }
    $dmConfig = new Configuration();
    $dmConfig->setDefaultDB(getenv('MONGODB_DATABASE', 'kubecampProgress'));
    $host = getenv('MONGODB_HOST');
    $port = getenv('MONGODB_PORT');
    $database = getenv('MONGODB_DATABASE');
    $url = "mongodb://$host:$port/$database";
    error_log("🍃 Connecting to $url");
    $connection = new Client($url, [], ['typeMap' => DocumentManager::CLIENT_TYPEMAP]);
    $dmConfig->setProxyDir(__DIR__ . '/../Proxies');
    $dmConfig->setProxyNamespace('Proxies');
    $dmConfig->setHydratorDir(__DIR__ . '/../Hydrators');
    $dmConfig->setHydratorNamespace('Hydrators');

    $annotationDriver = new AnnotationDriver(new AnnotationReader(), __DIR__ . '/../Models');
    $dmConfig->setMetadataDriverImpl($annotationDriver);


    spl_autoload_register($dmConfig->getProxyManagerConfiguration()->getProxyAutoloader());

    $documentManager = DocumentManager::create($connection, $dmConfig);

    self::$documentManager = $documentManager;
    return self::$documentManager;
  }

  static public function setDocumentManager($documentManager)
  {
    self::$documentManager = $documentManager;
  }
}
