package dev.arol.KubeCamp.Catalog.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import dev.arol.KubeCamp.Catalog.models.Course;
import dev.arol.KubeCamp.Catalog.repositories.CourseRepository;
import jakarta.persistence.EntityNotFoundException;

@Service
public class CourseService {
  private CourseRepository courseRepository;

  @Autowired
  public CourseService(CourseRepository courseRepository) {
    this.courseRepository = courseRepository;
  }

  public List<Course> getAllCourses() {
    return courseRepository.findAll();
  }

  public Course getCourseById(Long id) {
    return courseRepository.findById(id).orElseThrow(() -> new EntityNotFoundException("Course not found"));
  }

  public Course addCourse(Course course) {
    return courseRepository.save(course);
  }

  public Course updateCourse(Course course) {
    return courseRepository.save(course);
  }

}
