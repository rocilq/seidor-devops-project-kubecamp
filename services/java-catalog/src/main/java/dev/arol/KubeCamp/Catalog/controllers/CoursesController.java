package dev.arol.KubeCamp.Catalog.controllers;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import dev.arol.KubeCamp.Catalog.models.Course;
import dev.arol.KubeCamp.Catalog.models.Lesson;
import dev.arol.KubeCamp.Catalog.services.CourseService;
import dev.arol.KubeCamp.Catalog.services.LessonService;

@RestController
@RequestMapping("/api/courses")
public class CoursesController {

  private CourseService courseService;
  private LessonService lessonService;

  @Autowired
  public CoursesController(CourseService courseService, LessonService lessonService) {
    this.courseService = courseService;
    this.lessonService = lessonService;
  }

  @GetMapping
  public ResponseEntity<List<Course>> getAllCourses() {
    try {
      return ResponseEntity.ok(courseService.getAllCourses());
    } catch (Exception e) {
      return ResponseEntity.notFound().build();
    }
  }

  @PostMapping
  public ResponseEntity<Course> addCourse(@RequestBody Course course) {
    try {
      return new ResponseEntity<Course>(course, null, HttpStatus.CREATED);
    } catch (Exception e) {
      return ResponseEntity.badRequest().build();
    }
  }

  @GetMapping("/{id}")
  public ResponseEntity<Course> getCourse(@PathVariable Long id) {
    try {
      return ResponseEntity.ok(courseService.getCourseById(id));
    } catch (Exception e) {
      return ResponseEntity.notFound().build();
    }
  }

  @GetMapping("/{id}/lessons")
  public ResponseEntity<Object> getLessonsByCourseId(@PathVariable Long id) {
    try {
      return ResponseEntity.ok(lessonService.getAllLessonsByCourseId(id));
    } catch (Exception e) {
      Map<String, String> errorMap = new HashMap<>();
      if (e.getClass().getSimpleName().equals("EntityNotFoundException")) {
        errorMap.put("error", e.getMessage());
        return new ResponseEntity<>(errorMap, HttpStatus.NOT_FOUND);
      }
      return new ResponseEntity<>(errorMap, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }

  @PostMapping("/{id}/lessons")
  public ResponseEntity<Object> addLessonToCourse(@PathVariable Long id, @RequestBody Lesson lesson) {
    try {
      Course course = courseService.getCourseById(id);
      return ResponseEntity.ok(lessonService.addLessonToCourse(lesson, course));
    } catch (Exception e) {
      Map<String, String> errorMap = new HashMap<>();
      if (e.getClass().getSimpleName().equals("EntityNotFoundException")) {
        errorMap.put("error", e.getMessage());
        return new ResponseEntity<>(errorMap, HttpStatus.NOT_FOUND);
      }
      return new ResponseEntity<>(errorMap, HttpStatus.INTERNAL_SERVER_ERROR);
    }
  }

}
