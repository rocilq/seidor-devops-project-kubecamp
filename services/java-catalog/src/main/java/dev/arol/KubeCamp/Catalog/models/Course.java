package dev.arol.KubeCamp.Catalog.models;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonManagedReference;
import com.fasterxml.jackson.annotation.JsonProperty;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.OneToMany;

@Entity
public class Course {
  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Long id;
  private String name;
  private String description;

  @OneToMany(mappedBy = "course", fetch = FetchType.LAZY, cascade = CascadeType.ALL)
  @JsonManagedReference
  private List<Lesson> lessons;

  public Course() {
  }

  public Course(long id, String name, String description, List<Lesson> lessons) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.lessons = lessons;
  }

  public Course(@JsonProperty("name") String name, @JsonProperty("description") String description) {
    this.name = name;
    this.description = description;
  }

  public Long getId() {
    return this.id;
  }

  public String getName() {
    return this.name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getDescription() {
    return this.description;
  }

  public void setDesciption(String description) {
    this.description = description;
  }

  public List<Lesson> getLessons() {
    return this.lessons;
  }

  public void addLesson(Lesson lesson) {
    this.lessons.add(lesson);
  }

  public void removeLesson(Lesson lesson) {
    this.lessons.remove(lesson);
  }

}
