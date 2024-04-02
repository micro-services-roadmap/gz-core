package xerr

// OK 成功返回
const OK uint32 = 200

// region 全局错误码: (前3位代表业务,后三位代表具体功能)
const SERVER_INTERNAL_ERROR uint32 = 100_001
const REUQEST_PARAM_ERROR uint32 = 100_002
const TOKEN_EXPIRE_ERROR uint32 = 100_003
const TOKEN_GENERATE_ERROR uint32 = 100_004
const DB_ERROR uint32 = 100_005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100_006

// endregion
