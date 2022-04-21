package service

import (
	"github.com/smartschool/model/dto"
	"github.com/smartschool/repository"
)

func GetSemesterByFacultyID(facultyID uint) ([]dto.SemesterListElement, error) {
	semesterList, notFound, err := repository.QuerySemesterByFaculty(facultyID)
	if err != nil || notFound {
		return nil, err
	}

	return semesterList, nil
}
