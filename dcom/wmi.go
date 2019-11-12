/* WMI DCom Objtect Type definations */
package dcom

import (
	"local/msrpc/rpc"
)

type BSTR string   //FIXME
type LPCSTR string //FIXME
type LPWSTR string //FIXME

type REFGUID rpc.UUID
type WbemQueryFlagType uint32

const (
	WbemFLagDeep      WbemQueryFlagType = 0
	WbemFlagShallow                     = 1
	WbemFlagPrototype                   = 2
)

type WbemChangeFlagType uint32

const (
	WbemFlagCreateOrUpdate  WbemChangeFlagType = 0x00
	WbemFlagUpdateOnly                         = 0x01
	WbemFlagCreateOnly                         = 0x02
	WbemFlagUpdateSafeMode                     = 0x20
	WbemFlagUpdateForceMode                    = 0x40
)

type WbemConnectOptions uint32

const (
	WbemFlagConnectRepositoryOnly WbemConnectOptions = 0x40
	WbemFlagConnectProviders                         = 0x100
)

type WbemGenericFlagType uint32

const (
	WbemFlagReturnWbemComplete   WbemGenericFlagType = 0x0
	WbemFlagReturnImmediately                        = 0x10
	WbemFlagForwadOnly                               = 0x20
	WbemFlagNoErrorObject                            = 0x40
	WbemFlagSendStatus                               = 0x80
	WbemFlagEnsureLocatable                          = 0x100
	WbemFlagDirectRead                               = 0x200
	WbemMaskReserverFlags                            = 0x1F000
	WbemFlagUseAmendedQualifiers                     = 0x20000
	WbemFlagStrongValidation                         = 0x100000
)

type WbemTimeoutType uint32

const (
	WbemNoWait   WbemTimeoutType = 0x0
	WbemInfinite                 = 0xFFFFFFFF
)

type WbemBackupRestoreFlags uint32

const (
	WbemFlagBackupRestoreForceShutdown WbemBackupRestoreFlags = 1
)

type WbemStatus uint32

const (
	WbemSNoError                         = 0x00
	WbemSFalse                           = 0x01
	WbemSTimedout                        = 0x40004
	WbemSNewStyle                        = 0x400FF
	WbemSPartialResults                  = 0x40010
	WbemErrFailed                        = 0x80041001
	WbemErrNotFound                      = 0x80041002
	WbemErrAccessDenied                  = 0x80041003
	WbemErrProviderFailure               = 0x80041004
	WbemErrTyperrMismatch                = 0x80041005
	WbemErrOutOfMemory                   = 0x80041006
	WbemErrInvalidContext                = 0x80041007
	WbemErrInvalidParameter              = 0x80041008
	WbemErrNotAvailable                  = 0x80041009
	WbemErrCriticalError                 = 0x8004100a
	WbemErrNotSupported                  = 0x8004100C
	WbemErrProviderNotFound              = 0x80041011
	WbemErrInvalidProviderRegistration   = 0x80041012
	WbemErrProviderLoadFailure           = 0x80041013
	WbemErrInitializationFailure         = 0x80041014
	WbemErrTransportFailure              = 0x80041015
	WbemErrInvalidOperation              = 0x80041016
	WbemErrAlreadyExists                 = 0x80041019
	WbemErrUnexpected                    = 0x8004101d
	WbemErrIncompleterrClass             = 0x80041020
	WbemErrShuttingDown                  = 0x80041033
	ErrNotimpl                           = 0x80004001
	WbemErrInvalidSuperclass             = 0x8004100D
	WbemErrInvalidNamespace              = 0x8004100E
	WbemErrInvalidObject                 = 0x8004100F
	WbemErrInvalidClass                  = 0x80041010
	WbemErrInvalidQuery                  = 0x80041017
	WbemErrInvalidQueryType              = 0x80041018
	WbemErrProviderNotCapable            = 0x80041024
	WbemErrClassHasChildren              = 0x80041025
	WbemErrClassHasInstances             = 0x80041026
	WbemErrIllegalNull                   = 0x80041028
	WbemErrInvalidCimType                = 0x8004102D
	WbemErrInvalidMethod                 = 0x8004102E
	WbemErrInvalidMethodParameters       = 0x8004102F
	WbemErrInvalidProperty               = 0x80041031
	WbemErrCallCancelled                 = 0x80041032
	WbemErrInvalidObjectPath             = 0x8004103A
	WbemErrOutOfDiskSpace                = 0x8004103B
	WbemErrUnsupportedPutExtension       = 0x8004103D
	WbemErrQuotaViolation                = 0x8004106c
	WbemErrServerTooBusy                 = 0x80041045
	WbemErrMethodNotImplemented          = 0x80041055
	WbemErrMethodDisabled                = 0x80041056
	WbemErrUnparsablerrQuery             = 0x80041058
	WbemErrNotEventClass                 = 0x80041059
	WbemErrMissingGroupWithin            = 0x8004105A
	WbemErrMissingAggregationList        = 0x8004105B
	WbemErrPropertyNotAnObject           = 0x8004105c
	WbemErrAggregatingByObject           = 0x8004105d
	WbemErrBackupRestorerrWinmgmtRunning = 0x80041060
	WbemErrQueuerrOverflow               = 0x80041061
	WbemErrPrivilegerrNotHeld            = 0x80041062
	WbemErrInvalidOperator               = 0x80041063
	WbemErrCannotBerrAbstract            = 0x80041065
	WbemErrAmendedObject                 = 0x80041066
	WbemErrVetoPut                       = 0x8004107A
	WbemErrProviderSuspended             = 0x80041081
	WbemErrEncryptedConnectionRequired   = 0x80041087
	WbemErrProviderTimedOut              = 0x80041088
	WbemErrNoKey                         = 0x80041089
	WbemErrProviderDisabled              = 0x8004108a
	WbemErrRegistrationTooBroad          = 0x80042001
	WbemErrRegistrationTooPrecise        = 0x80042002
)

/*
type WbemRefrVersionNumber uint32

const (
	WbemRefresherVersion WbemRefrVersionNumber = 2
)

type WbemInstanceBlobType uint32

const (
	WbemBlobTypeAll   WbemInstanceBlobType = 2
	WbemBlobTypeError                      = 3
	WbemBlobTypeEnum                       = 4
)

type WbemRefreshedObject struct {
	RequestId  uint32
	BlobType   WbemInstanceBlobType
	BlobLength uint32
	BlobData   []byte
}

type WbemRefreshInfoRemote struct {
	Refresher  *IWbemRemoteRefresher
	WbemObject *IWbemClassObject
	Guid       GUID
}

type WbemRefreshInfoNonHiperf struct {
	wszNamespace []uint16
	WbemObject   *IWbemClassObject
}

type WbemRefreshType uint32

const (
	WbemRefreshTypeInvalid   = 0
	WbemRefreshTypeRemote    = 3
	WbemRefreshTypeNonHiperf = 6
)

type WbemRefreshInfo struct {
	Type uint32
	//Info Could be of these type HRESULT, WbemRefreshInfoNonHiperf, WbemRefreshInfoRemote
	CancelId uint32
}

type WbemRefresherId struct {
	MachineName []uint16
	ProcessId   uint16
	Id          GUID
}
*/

type WbemReconnectType uint32

const (
	WbemReconnectTypeObject WbemReconnectType = 0
	WbemReconnectTypeEnum                     = 1
	WbemReconnectTypeLast                     = 2
)

type WbemReconnectInfo struct {
	Type uint32
	Path []uint16
}

type WbemReconnectResuls struct {
	Id     uint32
	Result HRESULT
}

/* DComClasses in Wbem. */
//ClassId =  uuid(8BC3F05E-D86B-11d0-A075-00C04FB68820)
type WbemLevel1Login struct {
}

//ClassID = uuid(674B6698-EE92-11d0-AD71-00C04FD8FDFF)
type WbemContext struct {
}

//ClassId =  uuid(9A653086-174F-11d2-B5F9-00104B703EFD)
type WbemClassObject struct {
}

//ClassId = uuid(C49E32C6-BC8B-11d2-85D4-00105A1F8304)
type WbemBackupRestore struct {
}

/* Wbem ORpc Interfaces */
type ORpcInterface struct {
}

/*
IWbemClassObject
9A653086-174F-11d2-B5F9-00104B703EFD
*/
type IWbemClassObject struct { //TODO
}

/*
IWbemContext
9A653086-174F-11d2-B5F9-00104B703EFD
*/
type IWbemContext struct { //TODO
}

/* IWbemObjectSink
   uuid(7c857801-7381-11cf-884d-00aa004b2e24) */
type IWbemObjectSink struct { //TODO
}
type IWbemObjectSink_IndicateReq struct {
	ObjectCount uint32
	ObjectArray []IWbemClassObject
}

type IWbemObjectSink_SetStatusReq struct {
	ObjectCount uint32
	Hresult     HRESULT
	Param       string
	Object      IWbemClassObject
}

/* IEnumWbemClassObject
   uuid(027947e1-d731-11ce-a357-000000000001) */
type IEnumWbemClassObject struct { //TODO
}
type IEnumWbemClassObject_ResetReq struct {
}

type IEnumWbemClassObject_NextReq struct {
	Timeout uint32
	Count   uint32
}

type IEnumWbemClassObject_NextResp struct {
	Objects []IWbemClassObject
	Count   uint32
}

type IEnumWbemClassObject_NextAsyncReq struct {
	Count uint32
	PSink IWbemObjectSink
}

type IEnumWbemClassObject_CloneReq struct {
	Objects []IEnumWbemClassObject
}
type IEnumWbemClassObject_SkipReq struct {
	Timeout uint32
	Count   uint32
}

/*
 IWbemCallResult
 uuid(44aca675-e8fc-11d0-a07c-00c04fb68820)
*/
type IWbemCallResult struct { //TODO
}
type IWbemCallResult_GetResultObjectReq struct {
	TimeOut uint32
}
type IWbemCallResult_GetResultObjectResp struct {
	ResultObjects []IWbemClassObject
}

type IWbemCallResult_GetResultStringReq struct {
	Timeout uint32
}

type IWbemCallResult_GetResultStringResp struct {
	ResultString string
}

type IWbemCallResult_GetResultServicesReq struct {
	Timeout  uint32
	Services IWbemServices //PtrPtr
}

type IWbemCallResult_GetCallStatusReq struct {
	Timeout uint32
}
type IWbemCallResult_GetCallStatusResp struct {
	Status uint32
}

/*
IWbemServices
uuid(9556dc99-828c-11cf-a37e-00aa003240c7)
*/
type IWbemServices struct { //TODO
}
type IWbemServices_OpenNamespaceReq struct {
	Namespace BSTR
	Flags     uint32
	Ctx       IWbemContext
}

type IWbemServices_OpenNamespaceResp struct {
	Namespace IWbemServices
	Result    IWbemCallResult
}

type IWbemServices_CancelAsyncCallReq struct {
	Sink IWbemObjectSink
}

type IWbemServices_QueryObjectSinkReq struct {
	Flags uint32
}

type IWbemServices_QueryObjectSinkResp struct {
	ResponseHandler IWbemObjectSink
}

type IWbemServices_GetObjectReq struct {
	ObjectPath BSTR
	Flags      uint32
	Ctx        IWbemContext
}

type IWbemServices_GetObjectResp struct {
	Object     IWbemClassObject
	CallResult IWbemCallResult
}

type IWbemServices_GetObjectAsyncReq struct {
	ObjectPath      BSTR
	Flags           uint32
	Ctx             IWbemContext
	ResponseHandler IWbemObjectSink
}

type IWbemServices_PutClassReq struct {
	Object IWbemClassObject
	Flags  uint32
	Ctx    IWbemContext
}

type IWbemServices_PutClassResp struct {
	CallResult IWbemCallResult
}
type IWbemServices_PutClassAsyncReq struct {
	Object          IWbemClassObject
	Flags           uint32
	Context         IWbemContext
	ResponseHandler IWbemObjectSink
}

type IWbemServices_DeleteClassReq struct {
	Class   BSTR
	Flags   uint32
	Context IWbemContext
}
type IWbemServices_DeleteClassResp struct {
	CallResult IWbemCallResult
}

type IWbemServices_DeleteClassAsyncReq struct {
	Class           BSTR
	Flags           uint32
	Context         IWbemContext
	ResponseHandler IWbemObjectSink
}

type IWbemServices_CreateClassEnumReq struct {
	SupperClassName BSTR
	Flags           uint32
	Context         IWbemContext
	Enum            IEnumWbemClassObject
}

type IWbemServices_CreateClassEnumResp struct {
	CallResult IWbemCallResult
}

type IWbemServices_CreateClassEnumAsyncReq struct {
	SupperClassName BSTR
	Flags           uint32
	Context         IWbemContext
	Enum            IEnumWbemClassObject
	ResponseHandler IWbemObjectSink
}

type IWbemServices_PutInstanceReq struct {
	Instance IWbemClassObject
	Flags    uint32
	Context  IWbemContext
	Enum     IEnumWbemClassObject
}

type IWbemServices_PutInstanceResp struct {
	CallResult IWbemCallResult
}

type IWbemServices_PutInstanceAsyncReq struct {
	Instance        IWbemClassObject
	Flags           uint32
	Context         IWbemContext
	Enum            IEnumWbemClassObject
	ResponseHandler IWbemObjectSink
}
type IWbemServices_DeleteInstanceReq struct {
	ObjectPath BSTR
	Flags      uint32
	Context    IWbemContext
	Enum       IEnumWbemClassObject
}

type IWbemServices_DeleteInstanceResp struct {
	CallResult IWbemCallResult
}

type IWbemServices_DeleteInstanceAsyncReq struct {
	ObjectPath      BSTR
	Flags           uint32
	Context         IWbemContext
	Enum            IEnumWbemClassObject
	ResponseHandler IWbemObjectSink
}

type IWbemServices_CreateInstanceEnumReq struct {
	SupperClassName BSTR
	Flags           uint32
	Context         IWbemContext
	Enum            IEnumWbemClassObject
}

type IWbemServices_CreateInstanceEnumResp struct {
	CallResult IWbemCallResult
}

type IWbemServices_CreateInstanceEnumAsyncReq struct {
	SupperClassName BSTR
	Flags           uint32
	Context         IWbemContext
	Enum            IEnumWbemClassObject
	ResponseHandler IWbemObjectSink
}

type IWbemServices_ExecQueryReq struct {
	QueryLang BSTR
	Query     BSTR
	Flags     uint32
	Ctx       IWbemContext
}
type IWbemServices_ExecQueryResp struct {
	Enum IEnumWbemClassObject
}

type IWbemServices_ExecQueryAsyncReq struct {
	QueryLang       BSTR
	Query           BSTR
	Flags           uint32
	Ctx             IWbemContext
	ResponseHandler IWbemObjectSink
}
type IWbemServices_ExecNotificationQueryReq struct {
	QueryLang BSTR
	Query     BSTR
	Flags     uint32
	Ctx       IWbemContext
}
type IWbemServices_ExecNotificationQueryResp struct {
	Enum IEnumWbemClassObject
}

type IWbemServices_ExecNotificationQueryAsyncReq struct {
	QueryLang       BSTR
	Query           BSTR
	Flags           uint32
	Ctx             IWbemContext
	ResponseHandler IWbemObjectSink
}

type IWbemServices_ExecMethodReq struct {
	ObjectPath BSTR
	MethodName BSTR
	Flags      uint32
	Ctx        IWbemContext
	InParams   IWbemClassObject
}

type IWbemServices_ExecMethodResp struct {
	OutParams  IWbemClassObject
	CallResult IWbemClassObject
}
type IWbemServices_ExecMethodAsyncReq struct {
	ObjectPath      BSTR
	MethodName      BSTR
	Flags           uint32
	Ctx             IWbemContext
	InParams        IWbemClassObject
	ResponseHandler IWbemObjectSink
}

/*
   IWbemWCOSmartEnum
   uuid(423EC01E-2E35-11d2-B604-00104B703EFD)
*/
type IWbemWCOSmartEnum struct { //TODO
}
type IWbemWCOSmartEnum_NextReq struct {
	ProxyGuid REFGUID
	TimeOut   uint32
	Count     uint32
}
type IWbemWCOSmartEnum_NextResp struct {
	Returned uint32 //ptr
	BuffSize uint32 //ptr
	Buffer   []byte //size_is(BuffSize)
}

/*
IWbemFetchSmartEnum
uuid(1C1C45EE-4395-11d2-B60B-00104B703EFD)
*/
type IWbemFetchSmartEnum struct { //TODO
}
type IWbemFetchSmartEnum_GetSmartEnumReq struct {
}
type IWbemFetchSmartEnum_GetSmartEnumResp struct {
	SmartEnum IWbemWCOSmartEnum
}

/*
IWbemLoginClientID
uuid(d4781cd6-e5d3-44df-ad94-930efe48a887)
*/
type IWbemLoginClientId struct { //TODO
}

type IWbemLoginClientId_SetClientInfoReq struct {
	ClientMachine LPWSTR
	ClientProcId  uint32
	Reserver      uint32
}

/*
IWbemLevel1Login
uuid(F309AD18-D86A-11d0-A075-00C04FB68820)
*/
type IWbemLevel1Login struct { //TODO
}
type IWbemLevel1Login_EstablishPositionReq struct {
	Reserverd  []uint16 //string
	Reserverd2 DWORD    //Ptr
}

type IWbemLevel1Login_EstablishPositionResp struct {
	LocaleVersion DWORD //Ptr
}

type IWbemLevel1Login_RequestChallengeReq struct {
	Reserved1 []uint16 //string
	Reserved2 []uint16 //string
}

type IWbemLevel1Login_RequestChallengeResp struct {
	Reserved3 [16]byte
}

type IWbemLevel1Login_WBEMLoginReq struct {
	Reserved1 []uint16 //string
	Reserved2 [16]byte
	Reserved3 uint32
	Context   IWbemContext
}

type IWbemLevel1Login_WBEMLoginResp struct {
	Namespace IWbemServices
}

type IWbemLevel1Login_NTLMLoginReq struct {
	NetworkResource LPWSTR
	PreferredLocale LPWSTR
	Flags           uint32
	Context         IWbemContext
}

type IWbemLevel1Login_NTLMLoginResp struct {
	Namespace IWbemServices
}

/*
IWbemLoginHelper
uuid(541679AB-2E5F-11d3-B34E-00104BCC4B4A)
*/

type IWbemLoginHelper struct { //TODO
}
type IWbemLoginHelper_SetEventReq struct {
	EventToSet LPCSTR
}
