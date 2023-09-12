package dev.arol.KubeCamp.Catalog.repositories;

import java.util.List;

import org.springframework.data.jpa.repository.JpaRepository;

import dev.arol.KubeCamp.Catalog.models.Lesson;

public interface LessonRepository extends JpaRepository<Lesson, Long> {
  public List<Lesson> findAllByCourseId(Long courseId);
}
