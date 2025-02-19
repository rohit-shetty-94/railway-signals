package constants

const (
	// Success Messages
	MsgCreateSuccess = "Record created successfully"
	MsgReadSuccess   = "Record retrieved successfully"
	MsgsReadSuccess  = "Records retrieved successfully"
	MsgUpdateSuccess = "Record updated successfully"
	MsgDeleteSuccess = "Record deleted successfully"

	// Error Messages
	ErrRecordsRetrivedFail = "Retrieve records failed"
	ErrRecordUpdateFail    = "Update record failed"
	ErrRecordDeleteFail    = "Delete record failed"
	ErrRecordCreateFail    = "Create record failed"
	ErrRecordNotFound      = "Record not found"
	ErrInvalidInput        = "Invalid input data"

	// Track Signal Association
	TrackSignalAssociatedSuccess = "Track associated with signal"
	TrackSignalAssociatedRemoved = "Track-Signal association removed"
)
