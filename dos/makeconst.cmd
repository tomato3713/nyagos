pushd "%~dp0"
go run importconst_run.go -p dos ^
	CP_THREAD_ACP ^
	RESOURCE_CONNECTED ^
	RESOURCE_CONTEXT ^
	RESOURCE_GLOBALNET ^
	RESOURCE_REMEMBERED ^
	RESOURCETYPE_ANY ^
	RESOURCETYPE_DISK ^
	RESOURCETYPE_PRINT ^
	RESOURCEDISPLAYTYPE_NETWORK ^
	RESOURCEUSAGE_CONNECTABLE ^
	RESOURCEUSAGE_CONTAINER ^
	RESOURCEUSAGE_ATTACHED ^
	RESOURCEUSAGE_ALL ^
	ERROR_NOT_CONTAINER ^
	ERROR_INVALID_PARAMETER ^
	ERROR_NO_NETWORK ^
	ERROR_EXTENDED_ERROR ^
	ERROR_NO_MORE_ITEMS ^
	ERROR_MORE_DATA ^
	FSCTL_SET_REPARSE_POINT 
popd
