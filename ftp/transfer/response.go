package transfer

const (
	RestartMarkerReply         = "110 Restart marker reply."
	ReadyInMinutes             = "120 Service ready in nnn minutes."
	AlreadyOpen                = "125 Data connection already open; transfer starting."
	FileStatusOk               = "150 File status okay; about to open data connection."
	Ok                         = "200 Command okay."
	NotImplementedAtThisSite   = "202 Command not implemented, superfluous at this site."
	SystemStatus               = "211 System status, or system help reply."
	DirectoryStatus            = "212 Directory status."
	FileStatus                 = "213 File status."
	HelpMessage                = "214 Help message."
	NameSystemType             = "215 %s system type."
	ReadyForNewUser            = "220 Service ready for new user."
	ClosingControlConnection   = "221 Service closing control connection."
	ConnectionOpen             = "225 Data connection open; no transfer in progress."
	ClosingDataConnection      = "226 Closing data connection."
	EnteringPassiveMode        = "227 Entering Passive Mode (%s)."
	UserLoggedIn               = "230 User logged in, proceed."
	FileActionOk               = "250 Requested file action okay, completed."
	Created                    = "257 %q created."
	NeedPassword               = "331 User name okay, need password."
	NeedAccountForLogin        = "332 Need account for login."
	Peding                     = "350 Requested file action pending further information."
	LocalError                 = "421 Service not available, closing control connection."
	NotAvailable               = "425 Can't open data connection."
	CantOpenConnection         = "426 Connection closed; transfer aborted."
	ConnectionClosed           = "450 Requested file action not taken."
	LocalErrorrInProcessing    = "451 Requested action aborted: local error in processing."
	UnavailableFile            = "452 Requested action not taken."
	WrongCommand               = "500 Syntax error, command unrecognized."
	WrongArguments             = "501 Syntax error in parameters or arguments."
	NotImplemented             = "502 Command not implemented."
	BadSequence                = "503 Bad sequence of commands."
	NotImplementedForParameter = "504 Command not implemented for that parameter."
	NotLoggedIn                = "530 Not logged in."
	NeedAccountForStoringFiles = "532 Need account for storing files."
	ActionNotTaken             = "550 Requested action not taken."
	UnknownPageType            = "551 Requested action aborted: page type unknown."
	NotEnoughSpace             = "552 Requested file action aborted."
	DisallowedFileName         = "553 Requested action not taken."
)

type Response struct {
	Message string
	Data    string
	HasData bool
	Closing bool
}

func NewResponse(message string) Response {
	return Response{Message: message}
}

func (res *Response) SetData(data string) {
	res.HasData = true
	res.Data = data
}

func (res *Response) Close() {
	res.Closing = true
}

func (res *Response) String() string {
	return res.Message
}
