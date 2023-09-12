package dev.arol.KubeCamp.Catalog.repositories;

import org.springframework.data.jpa.repository.JpaRepository;

import dev.arol.KubeCamp.Catalog.models.Course;

public interface CourseRepository extends JpaRepository<Course, Long> {

}
