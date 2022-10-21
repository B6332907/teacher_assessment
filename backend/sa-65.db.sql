BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "teachers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"faculty_id"	integer,
	"level"	text,
	"name"	text,
	"email"	text,
	"graduate_faculty_level_id"	integer,
	"officer_id"	integer,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "students" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"college_year"	integer,
	"gpx"	integer,
	"faculty_id"	integer,
	"date_of_birth"	text,
	"phone"	text,
	"parent"	text,
	"teacher_id"	integer,
	"officer_id"	integer,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "teaching_durations" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"description"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "content_difficulty_levels" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"description"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "content_qualities" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"description"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "teacher_assessments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"timestamp"	datetime,
	"student_id"	integer,
	"teacher_id"	integer,
	"teaching_duration_id"	integer,
	"content_difficulty_level_id"	integer,
	"content_quality_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_teacher_assessments_teacher" FOREIGN KEY("teacher_id") REFERENCES "teachers"("id"),
	CONSTRAINT "fk_teacher_assessments_teaching_duration" FOREIGN KEY("teaching_duration_id") REFERENCES "teaching_durations"("id"),
	CONSTRAINT "fk_teacher_assessments_content_quality" FOREIGN KEY("content_quality_id") REFERENCES "content_qualities"("id"),
	CONSTRAINT "fk_teacher_assessments_content_difficulty_level" FOREIGN KEY("content_difficulty_level_id") REFERENCES "content_difficulty_levels"("id"),
	CONSTRAINT "fk_teacher_assessments_student" FOREIGN KEY("student_id") REFERENCES "students"("id")
);
CREATE INDEX IF NOT EXISTS "idx_teachers_deleted_at" ON "teachers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_students_deleted_at" ON "students" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_teaching_durations_deleted_at" ON "teaching_durations" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_content_difficulty_levels_deleted_at" ON "content_difficulty_levels" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_content_qualities_deleted_at" ON "content_qualities" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_teacher_assessments_deleted_at" ON "teacher_assessments" (
	"deleted_at"
);
COMMIT;
