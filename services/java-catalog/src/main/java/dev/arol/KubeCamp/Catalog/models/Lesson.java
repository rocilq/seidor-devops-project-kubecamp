package dev.arol.KubeCamp.Catalog.models;

import com.fasterxml.jackson.annotation.JsonBackReference;
import com.fasterxml.jackson.annotation.JsonProperty;

import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;

@Entity
public class Lesson {
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Long id;
  private String name;
  private String description;

  @ManyToOne(fetch = FetchType.LAZY)
  @JoinColumn(name = "course_id")
  @JsonBackReference
  private Course course;

  public Lesson() {
  }

  public Lesson(long id, String name, String description, Course course) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.course = course;
  }

  public Lesson(@JsonProperty("name") String name, @JsonProperty("description") String description) {
    this.name = name;
    this.description = description;
  }

  public Long getId() {
    return this.id;
  }

  public String getName() {
    return this.name;
  }

  public void setName(String title) {
    this.name = title;
  }

  public String getDescription() {
    return this.description;
  }

  public void setDesciption(String description) {
    this.description = description;
  }

  public Course getCourse() {
    return this.course;
  }

  public void setCourse(Course course) {
    this.course = course;
  }
}
