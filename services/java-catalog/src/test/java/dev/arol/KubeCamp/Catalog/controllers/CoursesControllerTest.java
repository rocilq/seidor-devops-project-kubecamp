package dev.arol.KubeCamp.Catalog.controllers;

import static org.mockito.Mockito.when;

import java.util.List;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.test.web.servlet.MockMvc;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

import com.fasterxml.jackson.databind.ObjectMapper;

import dev.arol.KubeCamp.Catalog.models.Course;
import dev.arol.KubeCamp.Catalog.models.Lesson;
import dev.arol.KubeCamp.Catalog.services.CourseService;
import dev.arol.KubeCamp.Catalog.services.LessonService;
import jakarta.persistence.EntityNotFoundException;

@WebMvcTest(CoursesController.class)
public class CoursesControllerTest {
  @Autowired
  private MockMvc mockMvc;

  @Autowired
  private ObjectMapper objectMapper;

  @MockBean
  private CourseService courseService;

  @MockBean
  private LessonService lessonsService;

  @Test
  public void testGetAllCourses() throws Exception {
    Lesson lesson1 = new Lesson(1L, "Lesson 1", "Description 1", null);
    Course course1 = new Course(1L, "Course 1", "Description 1", List.of(lesson1));
    when(courseService.getAllCourses()).thenReturn(List.of(course1));

    mockMvc.perform(get("/api/courses"))
        .andExpect(status().isOk())
        .andExpect(content().json(objectMapper.writeValueAsString(List.of(course1))));
  }

  @Test
  public void testGetCourseById() throws Exception {
    Lesson lesson1 = new Lesson(1L, "Lesson 1", "Description 1", null);
    Course course1 = new Course(1L, "Course 1", "Description 1", List.of(lesson1));
    when(courseService.getCourseById(1L)).thenReturn(course1);

    mockMvc.perform(get("/api/courses/1"))
        .andExpect(status().isOk())
        .andExpect(content().json(objectMapper.writeValueAsString(course1)));
  }

  @Test
  public void testGetCourseByIdNotFound() throws Exception {
    when(courseService.getCourseById(1L)).thenThrow(new EntityNotFoundException("Course not found"));

    mockMvc.perform(get("/api/courses/1"))
        .andExpect(status().isNotFound());
  }

  @Test
  public void testCreateCourse() throws Exception {
    Lesson lesson1 = new Lesson(1L, "Lesson 1", "Description 1", null);
    Course course1 = new Course(1L, "Course 1", "Description 1", List.of(lesson1));
    when(courseService.addCourse(course1)).thenReturn(course1);

    mockMvc.perform(post("/api/courses")
        .contentType("application/json")
        .content(objectMapper.writeValueAsString(course1)))
        .andExpect(status().isCreated())
        .andExpect(content().json(objectMapper.writeValueAsString(course1)));
  }
}