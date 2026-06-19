package domain

type StudentCourse struct {
	StudentCourseId int64 `json:"student_course_id"`
	StudentId       int64 `json:"student_id"`
	CourseId        int64 `json:"course_id"`
}
