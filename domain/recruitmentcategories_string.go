// Code generated by "stringer -type=RecruitmentCategories"; DO NOT EDIT.

package domain

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PROJECT-0]
	_ = x[STUDY-1]
}

const _RecruitmentCategories_name = "PROJECTSTUDY"

var _RecruitmentCategories_index = [...]uint8{0, 7, 12}

func (i RecruitmentCategories) String() string {
	if i < 0 || i >= RecruitmentCategories(len(_RecruitmentCategories_index)-1) {
		return "RecruitmentCategories(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RecruitmentCategories_name[_RecruitmentCategories_index[i]:_RecruitmentCategories_index[i+1]]
}
