package dev.arol.KubeCamp.Catalog.config;

import java.util.Arrays;
import java.util.List;

import org.slf4j.Logger;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import dev.arol.KubeCamp.Catalog.models.Course;
import dev.arol.KubeCamp.Catalog.models.Lesson;
import dev.arol.KubeCamp.Catalog.repositories.CourseRepository;
import dev.arol.KubeCamp.Catalog.repositories.LessonRepository;

@Component
public class DatabaseLoader implements CommandLineRunner {
  private final CourseRepository courseRepository;
  private final LessonRepository lessonRepository;
  private static final Logger logger = org.slf4j.LoggerFactory.getLogger(DatabaseLoader.class);

  @Autowired
  public DatabaseLoader(CourseRepository courseRepository, LessonRepository lessonRepository) {
    this.courseRepository = courseRepository;
    this.lessonRepository = lessonRepository;
  }

  @Override
  public void run(String... args) throws Exception {
    logger.info("📀 Loading data into database…");

    List<Course> courses = courseRepository.findAll();
    if (courses.size() > 0) {
      logger.info("⏩ Database already loaded. Skiping…");
      return;
    }

    Course course1 = new Course("Kubernetes 101", "Learn the basics of Kubernetes");
    Course course2 = new Course("Kubernetes Master", "Learn the advanced topics of Kubernetes");
    courseRepository.saveAll(Arrays.asList(course1, course2));
    logger.info("✅ Courses created");

    Lesson lesson1 = new Lesson("Introduction to Kubernetes", "Learn the basics of Kubernetes");
    Lesson lesson2 = new Lesson("Kubernetes Architecture", "Learn the architecture of Kubernetes");
    lesson1.setCourse(course1);
    lesson2.setCourse(course1);

    Lesson lesson3 = new Lesson("Kubernetes Networking", "Learn the networking of Kubernetes");
    Lesson lesson4 = new Lesson("Kubernetes Storage", "Learn the storage of Kubernetes");
    lesson3.setCourse(course2);
    lesson4.setCourse(course2);

    lessonRepository.saveAll(Arrays.asList(lesson1, lesson2, lesson3, lesson4));
    logger.info("✅ Lessons created");
    logger.info("✅ Finished loading data into database");
  }
}
