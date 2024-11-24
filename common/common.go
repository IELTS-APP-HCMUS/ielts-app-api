package common

// Define New Table Name here
const (
	POSTGRES_TABLE_NAME_USERS           = "public.users"
	POSTGRES_TABLE_NAME_ROLES           = "public.roles"
	POSTGRES_TABLE_NAME_TARGETS         = "public.student_target"
	POSTGRES_TABLE_NAME_QUIZ            = "public.quiz"
	POSTGRES_TABLE_NAME_PART            = "public.part"
	POSTGRES_TABLE_NAME_QUESTION        = "public.question"
	POSTGRES_TABLE_NAME_TAG_SEARCH      = "public.tag_search"
	POSTGRES_TABLE_NAME_TAG_POSITION    = "public.tag_search_position"
	POSTGRES_TABLE_NAME_QUIZ_TAG_SEARCH = "public.quiz_tag_search"
	POSTGRES_TABLE_NAME_QUIZ_SKILL      = "public.type"
	POSTGRES_TABLE_NAME_QUIZ_PART       = "public.quiz_part"
)

// Define New Const variable here for each service

const (
	ROLE_END_USER        = "end_user"
	ROLE_END_USER_UUID   = "da0e07d4-ce51-4784-a5a9-a018434adf8e"
	USER_PROVIDER_GOOGLE = "google"
)

// Define other common variable here
const (
	QUESTION_TYPE_SINGLE_RADIO      = "SINGLE-RADIO"
	QUESTION_TYPE_SINGLE_SELECTION  = "SINGLE-SELECTION"
	QUESTION_TYPE_FILL_IN_THE_BLANK = "FILL-IN-THE-BLANK"
	QUESTION_TYPE_MULTIPLE          = "MULTIPLE"
)
