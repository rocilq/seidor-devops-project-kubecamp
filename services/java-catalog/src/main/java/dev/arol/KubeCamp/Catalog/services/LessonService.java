package dev.arol.KubeCamp.Catalog.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import dev.arol.KubeCamp.Catalog.models.Course;
import dev.arol.KubeCamp.Catalog.models.Lesson;
import dev.arol.KubeCamp.Catalog.repositories.LessonRepository;
import jakarta.persistence.EntityNotFoundException;

@Service
public class LessonService {
  private LessonRepository lessonRepository;

  @Autowired
  public LessonService(LessonRepository lessonRepository) {
    this.lessonRepository = lessonRepository;
  }

  public List<Lesson> getAllLessonsByCourseId(Long courseId) {
    // TODO: implement findAllByCourseId from interface
    return lessonRepository.findAllByCourseId(courseId);
  }

  public Lesson getLessonById(Long id) {
    return lessonRepository.findById(id).orElseThrow(() -> new EntityNotFoundException("Lesson not found"));
  }

  public Lesson addLessonToCourse(Lesson lesson, Course course) {
    lesson.setCourse(course);
    return lessonRepository.save(lesson);
  }

}
