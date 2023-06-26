package main

import (
	"context"
	management "example_hertz_kitex/kitex/kitex_gen/student/management"
)

// StudentManagementImpl implements the last service interface defined in the IDL.
type StudentManagementImpl struct{}

// For QueryStudentResponse, each student contains these info
type StudentInfo struct {
	Num string
	Name string
	Gender string
}

// Stores 5 student info first
var StudentData = make(map[string]StudentInfo, 5)

// QueryStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) QueryStudent(ctx context.Context, req *management.QueryStudentRequest) (resp *management.QueryStudentResponse, err error) {
	// TODO: Your code here...
	// When accessing a key in a map, can return two values,
	// the first containing the value, the second being true/false
	stu, exist := StudentData[req.Num]

	if !exist {
		return &management.QueryStudentResponse{
			Exist: false,
		}, nil
	}

	resp = &management.QueryStudentResponse{
		Exist: true,
		Num: stu.Num,
		Name: stu.Name,
		Gender: stu.Gender,
	}

	return resp, nil
}

// InsertStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) InsertStudent(ctx context.Context, req *management.InsertStudentRequest) (resp *management.InsertStudentResponse, err error) {
	// TODO: Your code here...
	// We don't need the student, just need to check if he exists
	_, exist := StudentData[req.Num]
	
	// Unsuccessful response
	if exist {
		return &management.InsertStudentResponse{
			Ok: false,
			Msg: "There is a student at that number already.",
		}, nil
	}

	// Insert if does not exist
	StudentData[req.Num] = StudentInfo{
		Num: req.Num,
		Name: req.Name,
		Gender: req.Gender,
	}

	// Successful response
	resp = &management.InsertStudentResponse{
		Ok: true,
		Msg: "Student successfully inserted.",
	}

	return resp, nil
}
