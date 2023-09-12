package dev.arol.KubeCamp.Catalog.controllers;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import dev.arol.KubeCamp.Catalog.models.Lesson;
import dev.arol.KubeCamp.Catalog.services.LessonService;

@RestController
@RequestMapping("/api/lessons")
public class LessonsController {
  private LessonService lessonService;

  @Autowired
  public LessonsController(LessonService lessonService) {
    this.lessonService = lessonService;
  }

  @GetMapping("/{id}")
  public ResponseEntity<Lesson> getLesson(@PathVariable Long id) {
    try {
      return ResponseEntity.ok(lessonService.getLessonById(id));
    } catch (Exception e) {
      return ResponseEntity.notFound().build();
    }
  }

}
