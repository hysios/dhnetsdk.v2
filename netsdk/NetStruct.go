// NetSDK  NetSDKStruct.go

package netsdk

import (
	_ "fmt"
	_ "unsafe"
)

type LLONG = int64
type DWORD = uint32
type WORD = uint16
type BYTE = byte
type BOOL = int32

type NETSDK_INIT_PARAM struct {
	ST_bReserved  [1024]byte "保留字节"
	ST_nThreadNum int        "指定NetSDK常规网络处理线程数, 当值为0时, 使用内部默认值"
}

type EVENT_IVS int32

const (
	EVENT_IVS_ALL                  EVENT_IVS = 0x00000001 // 订阅所有事件
	EVENT_IVS_CROSSLINEDETECTION             = 0x00000002 // 警戒线事件(对应 DEV_EVENT_CROSSLINE_INFO)
	EVENT_IVS_CROSSREGIONDETECTION           = 0x00000003 // 警戒区事件(对应 DEV_EVENT_CROSSREGION_INFO)
	EVENT_IVS_TRAFFICJUNCTION                = 0x00000017 // 交通路口事件----老规则(对应 DEV_EVENT_TRAFFICJUNCTION_INFO)
	EVENT_IVS_TRAFFIC_PARKING                = 0x00000108
	EVENT_ALARM_MOTIONDETECT                 = 0x0000011C // 视频移动侦测事件(对应 DEV_EVENT_ALARM_INFO)
)

type EM_LOGIN_SPAC_CAP_TYPE int32

const (
	EM_LOGIN_SPEC_CAP_TCP            EM_LOGIN_SPAC_CAP_TYPE = iota // TCP登陆, 默认方式
	EM_LOGIN_SPEC_CAP_ANY                                          // 无条件登陆
	EM_LOGIN_SPEC_CAP_SERVER_CONN                                  // 主动注册的登入
	EM_LOGIN_SPEC_CAP_MULTICAST                                    // 组播登陆
	EM_LOGIN_SPEC_CAP_UDP                                          // UDP方式下的登入
	EM_LOGIN_SPEC_CAP_MAIN_CONN_ONLY                               // 只建主连接下的登入
	EM_LOGIN_SPEC_CAP_SSL                                          // SSL加密方式登陆
)

type EM_TRAFFICCAR_MOVE_DIRECTION int32 // 交通车辆行驶方向类型
const (
	EM_TRAFFICCAR_MOVE_DIRECTION_UNKNOWN     EM_TRAFFICCAR_MOVE_DIRECTION = iota // 未知的
	EM_TRAFFICCAR_MOVE_DIRECTION_STRAIGHT                                        // 直行
	EM_TRAFFICCAR_MOVE_DIRECTION_TURN_LEFT                                       // 左转
	EM_TRAFFICCAR_MOVE_DIRECTION_TURN_RIGHT                                      // 右转
	EM_TRAFFICCAR_MOVE_DIRECTION_TURN_AROUND                                     // 掉头
)

type EM_TRAFFICCAR_CAR_TYPE int32 // 车辆类型
const (
	EM_TRAFFICCAR_CAR_TYPE_UNKNOWN        EM_TRAFFICCAR_CAR_TYPE = iota // 未知的
	EM_TRAFFICCAR_CAR_TYPE_TRUST_CAR                                    // 白名单车辆
	EM_TRAFFICCAR_CAR_TYPE_SUSPICIOUS_CAR                               //  黑名单车辆
	EM_TRAFFICCAR_CAR_TYPE_NORMAL_CAR                                   // 非白名单且非黑名单车辆
)

type EM_TRAFFICCAR_LANE_TYPE int32 // 车道类型
// 车道类型
const (
	EM_TRAFFICCAR_LANE_TYPE_UNKNOWN    EM_TRAFFICCAR_LANE_TYPE = iota // 未知
	EM_TRAFFICCAR_LANE_TYPE_NORMAL                                    // 普通车道
	EM_TRAFFICCAR_LANE_TYPE_NONMOTOR                                  // 非机动车车道
	EM_TRAFFICCAR_LANE_TYPE_LIGHT_DUTY                                // 小型车车道
	EM_TRAFFICCAR_LANE_TYPE_BUS                                       // 公交车车道
	EM_TRAFFICCAR_LANE_TYPE_EMERGENCY                                 // 应急车道
	EM_TRAFFICCAR_LANE_TYPE_DANGEROUS                                 // 危险品车道
)

type EM_PLATE_COLOR_TYPE int32 // 车牌颜色
const (
	EM_PLATE_COLOR_UNKNOWN      EM_PLATE_COLOR_TYPE = iota // 未知 "Unknown"
	EM_PLATE_COLOR_OTHER                                   // 其他颜色 "Other"
	EM_PLATE_COLOR_BLUE                                    // 蓝色 "Blue"
	EM_PLATE_COLOR_YELLOW                                  // 黄色 "Yellow"
	EM_PLATE_COLOR_WHITE                                   // 白色 "White"
	EM_PLATE_COLOR_BLACK                                   // 黑色 "Black"
	EM_PLATE_COLOR_RED                                     // 红色 "Red"
	EM_PLATE_COLOR_GREEN                                   // 绿色 "Green"
	EM_PLATE_COLOR_SHADOW_GREEN                            // 渐变绿 "ShadowGreen"
	EM_PLATE_COLOR_YELLOW_GREEN                            // 黄绿双拼 "YellowGreen""
)

type EM_FEATURE_VERSION int32 // 特征值版本类型
const (
	EM_FEATURE_VERSION_UNKNOWN                        EM_PLATE_COLOR_TYPE = iota     // 未知
	EM_FEATURE_VERSION_FACE_LARGE_1_01_001            EM_FEATURE_VERSION  = iota + 1 // 人脸，大模型，1.01.001
	EM_FEATURE_VERSION_FACE_LARGE_1_02_001                                           // 人脸，大模型，1.02.001
	EM_FEATURE_VERSION_FACE_LARGE_1_03_001                                           // 人脸，大模型，1.03.001
	EM_FEATURE_VERSION_FACE_LARGE_1_04_001                                           // 人脸，大模型，1.04.001
	EM_FEATURE_VERSION_FACE_MIDDLE_1_01_002           = iota + 31                    // 人脸，中模型，1.01.002
	EM_FEATURE_VERSION_FACE_MIDDLE_1_02_002                                          // 人脸，中模型，1.02.002
	EM_FEATURE_VERSION_FACE_MIDDLE_1_03_002                                          // 人脸，中模型，1.03.002
	EM_FEATURE_VERSION_FACE_MIDDLE_1_04_002                                          // 人脸，中模型，1.04.002
	EM_FEATURE_VERSION_FACE_SMALL_1_01_003            = iota + 61                    // 人脸，小模型，1.01.003
	EM_FEATURE_VERSION_FACE_SMALL_1_02_003                                           // 人脸，小模型，1.02.003
	EM_FEATURE_VERSION_HUMAN_NONMOTOR                 = iota + 91                    // 人和非机动车，全局无版本号
	EM_FEATURE_VERSION_HUMAN_NONMOTOR_FLOAT_1_00_01                                  // 人和非机动车，全局浮点，1.00.01
	EM_FEATURE_VERSION_HUMAN_NONMOTOR_HASH_1_00_01                                   // 人和非机动车，全局哈希，1.00.01
	EM_FEATURE_VERSION_HUMAN_NONMOTOR_FLOAT_1_01_00                                  // 人和非机动车，全局浮点，1.01.00
	EM_FEATURE_VERSION_HUMAN_NONMOTOR_HASH_1_01_00                                   // 人和非机动车，全局哈希，1.01.00
	EM_FEATURE_VERSION_TRAFFIC                        = iota + 121                   // 机动车，全局无版本号
	EM_FEATURE_VERSION_TRAFFIC_FLOAT                                                 // 机动车，全局浮点版本号0
	EM_FEATURE_VERSION_TRAFFIC_FLOAT_1_00_01                                         // 机动车，全局浮点版本号1.00.01
	EM_FEATURE_VERSION_TRAFFIC_HASH_1_00_01                                          // 机动车，全局哈希版本号1.00.01
	EM_FEATURE_VERSION_TRAFFIC_FLOAT_1_00_02                                         // 机动车，全局浮点版本号1.00.02
	EM_FEATURE_VERSION_TRAFFIC_HASH_1_00_02                                          // 机动车，全局哈希版本号1.00.02
	EM_FEATURE_VERSION_SHANGTANG_FACE_1_5_0           = iota + 151                   // 商汤，人脸，1.5.0
	EM_FEATURE_VERSION_SHANGTANG_FACE_1_8_1                                          // 商汤，人脸，1.8.1
	EM_FEATURE_VERSION_SHANGTANG_FACE_2_1_3                                          // 商汤，人脸，2.1.3
	EM_FEATURE_VERSION_SHANGTANG_FACE_2_39_6                                         // 商汤，人脸，2.39.6
	EM_FEATURE_VERSION_SHANGTANG_FACE_2_39_7                                         // 商汤，人脸，2.39.7
	EM_FEATURE_VERSION_SHANGTANG_FACE_2_39_8                                         // 商汤，人脸，2.39.8
	EM_FEATURE_VERSION_SHANGTANG_FACE_239                                            // 商汤，人脸，239
	EM_FEATURE_VERSION_SHANGTANG_FACE_242                                            // 商汤，人脸，242
	EM_FEATURE_VERSION_SHANGTANG_FACE_244                                            // 商汤，人脸，244
	EM_FEATURE_VERSION_SHANGTANG_FACE_245                                            // 商汤，人脸，245
	EM_FEATURE_VERSION_SHENMO_HUMAN_TRAFFIC_NON_2_4_2 = iota + 181                   // 深瞐，人脸/机动车/非机动车，2.4.2
	EM_FEATURE_VERSION_SHENMO_HUMAN_TRAFFIC_NON_2_5_7                                // 深瞐，人脸/机动车/非机动车，2.5.7
)

type EM_OBJECT_COLOR_TYPE int32 // 颜色类型
const (
	EM_OBJECT_COLOR_TYPE_UNKNOWN        EM_OBJECT_COLOR_TYPE = iota // 未知
	EM_OBJECT_COLOR_TYPE_WHITE                                      // 白色
	EM_OBJECT_COLOR_TYPE_ORANGE                                     // 橙色
	EM_OBJECT_COLOR_TYPE_PINK                                       // 粉色
	EM_OBJECT_COLOR_TYPE_BLACK                                      // 黑色
	EM_OBJECT_COLOR_TYPE_RED                                        // 红色
	EM_OBJECT_COLOR_TYPE_YELLOW                                     // 黄色
	EM_OBJECT_COLOR_TYPE_GRAY                                       // 灰色
	EM_OBJECT_COLOR_TYPE_BLUE                                       // 蓝色
	EM_OBJECT_COLOR_TYPE_GREEN                                      // 绿色
	EM_OBJECT_COLOR_TYPE_PURPLE                                     // 紫色
	EM_OBJECT_COLOR_TYPE_BROWN                                      // 棕色
	EM_OBJECT_COLOR_TYPE_SLIVER                                     // 银色
	EM_OBJECT_COLOR_TYPE_DARKVIOLET                                 // 暗紫罗兰色
	EM_OBJECT_COLOR_TYPE_MAROON                                     // 栗色
	EM_OBJECT_COLOR_TYPE_DIMGRAY                                    // 暗灰色
	EM_OBJECT_COLOR_TYPE_WHITESMOKE                                 // 白烟色
	EM_OBJECT_COLOR_TYPE_DARKORANGE                                 // 深橙色
	EM_OBJECT_COLOR_TYPE_MISTYROSE                                  // 浅玫瑰色
	EM_OBJECT_COLOR_TYPE_TOMATO                                     // 番茄红色
	EM_OBJECT_COLOR_TYPE_OLIVE                                      // 橄榄色
	EM_OBJECT_COLOR_TYPE_GOLD                                       // 金色
	EM_OBJECT_COLOR_TYPE_DARKOLIVEGREEN                             // 暗橄榄绿色
	EM_OBJECT_COLOR_TYPE_CHARTREUSE                                 // 黄绿色
	EM_OBJECT_COLOR_TYPE_GREENYELLOW                                // 绿黄色
	EM_OBJECT_COLOR_TYPE_FORESTGREEN                                // 森林绿色
	EM_OBJECT_COLOR_TYPE_SEAGREEN                                   // 海洋绿色
	EM_OBJECT_COLOR_TYPE_DEEPSKYBLUE                                // 深天蓝色
	EM_OBJECT_COLOR_TYPE_CYAN                                       // 青色
	EM_OBJECT_COLOR_TYPE_OTHER                                      // 无法识别
)

type EM_STRABISMUS_TYPE int32 // 斜视状态
const (
	EM_STRABISMUS_UNKNOWN EM_STRABISMUS_TYPE = iota // 未知
	EM_STRABISMUS_NORMAL                            // 正常
	EM_STRABISMUS_YES                               // 斜视
)

type EM_COMPLEXION_TYPE int32 // 肤色
const (
	EM_COMPLEXION_NODISTI EM_COMPLEXION_TYPE = iota // 未识别
	EM_COMPLEXION_YELLOW                            // 黄
	EM_COMPLEXION_BLACK                             // 黑
	EM_COMPLEXION_WHITE                             // 白
)

type EM_HAS_GLASS int32 // 是否带眼镜
const (
	EM_HAS_GLASS_UNKNOWN EM_HAS_GLASS = iota // 未知
	EM_HAS_GLASS_NO                          // 未戴
	EM_HAS_GLASS_NORMAL                      // 戴普通眼镜
	EM_HAS_GLASS_SUN                         // 戴太阳眼镜
	EM_HAS_GLASS_BLACK                       // 戴黑框眼镜
)

type EM_NONMOTOR_OBJECT_STATUS int32 // 事件/物体状态
const (
	EM_NONMOTOR_OBJECT_STATUS_UNKNOWN EM_NONMOTOR_OBJECT_STATUS = iota // 未识别
	EM_NONMOTOR_OBJECT_STATUS_NO                                       // 否
	EM_NONMOTOR_OBJECT_STATUS_YES                                      // 是
)

type EM_EMOTION_TYPE int32 // 表情
const (
	EM_EMOTION_TYPE_UNKNOWN  EM_EMOTION_TYPE = iota // 未知
	EM_EMOTION_TYPE_NORMAL                          // 普通/正常
	EM_EMOTION_TYPE_SMILE                           // 微笑
	EM_EMOTION_TYPE_ANGER                           // 愤怒
	EM_EMOTION_TYPE_SADNESS                         // 悲伤
	EM_EMOTION_TYPE_DISGUST                         // 厌恶
	EM_EMOTION_TYPE_FEAR                            // 害怕
	EM_EMOTION_TYPE_SURPRISE                        // 惊讶
	EM_EMOTION_TYPE_NEUTRAL                         // 正常
	EM_EMOTION_TYPE_LAUGH                           // 大笑
	EM_EMOTION_TYPE_HAPPY                           // 高兴
	EM_EMOTION_TYPE_CONFUSED                        // 困惑
	EM_EMOTION_TYPE_SCREAM                          // 尖叫
	EM_EMOTION_TYPE_CALMNESS                        // 平静
)

type EM_CLOTHES_TYPE int32

const (
	EM_CLOTHES_TYPE_UNKNOWN      EM_CLOTHES_TYPE = iota //未知
	EM_CLOTHES_TYPE_LONG_SLEEVE                         //长袖
	EM_CLOTHES_TYPE_SHORT_SLEEVE                        //短袖
	EM_CLOTHES_TYPE_TROUSERS                            //长裤
	EM_CLOTHES_TYPE_SHORTS                              //短裤
	EM_CLOTHES_TYPE_SKIRT                               //裙子
	EM_CLOTHES_TYPE_WAISTCOAT                           //背心
	EM_CLOTHES_TYPE_MINIPANTS                           //超短裤
	EM_CLOTHES_TYPE_MINISKIRT                           //超短裙
)

type EM_CATEGORY_NONMOTOR_TYPE int32 // 非机动车子类型
const (
	EM_CATEGORY_NONMOTOR_TYPE_UNKNOWN                     EM_CATEGORY_NONMOTOR_TYPE = iota // 未知
	EM_CATEGORY_NONMOTOR_TYPE_TRICYCLE                                                     // "Tricycle" 三轮车
	EM_CATEGORY_NONMOTOR_TYPE_MOTORCYCLE                                                   // "Motorcycle" 摩托车
	EM_CATEGORY_NONMOTOR_TYPE_NON_MOTOR                                                    // "Non-Motor" 非机动车
	EM_CATEGORY_NONMOTOR_TYPE_BICYCLE                                                      // "Bicycle" 自行车
	EM_CATEGORY_NONMOTOR_TYPE_DUALTRIWHEELMOTORCYCLE                                       // "DualTriWheelMotorcycle"两、三轮摩托车
	EM_CATEGORY_NONMOTOR_TYPE_LIGHTMOTORCYCLE                                              // "LightMotorcycle" 轻便摩托车
	EM_CATEGORY_NONMOTOR_TYPE_EMBASSYMOTORCYCLE                                            // "EmbassyMotorcycle "使馆摩托车
	EM_CATEGORY_NONMOTOR_TYPE_MARGINALMOTORCYCLE                                           // "MarginalMotorcycle " 领馆摩托车
	EM_CATEGORY_NONMOTOR_TYPE_AREAOUTMOTORCYCLE                                            // "AreaoutMotorcycle " 境外摩托车
	EM_CATEGORY_NONMOTOR_TYPE_FOREIGNMOTORCYCLE                                            // "ForeignMotorcycle " 外籍摩托车
	EM_CATEGORY_NONMOTOR_TYPE_TRIALMOTORCYCLE                                              // "TrialMotorcycle " 试验摩托车
	EM_CATEGORY_NONMOTOR_TYPE_COACHMOTORCYCLE                                              // "CoachMotorcycle " 教练摩托车
	EM_CATEGORY_NONMOTOR_TYPE_PASSERBY                                                     // "Passerby" 行人
	EM_CATEGORY_NONMOTOR_TYPE_VANTRICYCLE                                                  // "VanTricycle" 厢式三轮车
	EM_CATEGORY_NONMOTOR_TYPE_MANNEDCONVERTIBLETRICYCLE                                    // "MannedConvertibleTricycle" 载人敞篷三轮车
	EM_CATEGORY_NONMOTOR_TYPE_NOMANNEDCONVERTIBLETRICYCLE                                  // "NoMannedConvertibleTricycle" 不载人敞篷三轮车
)

type EM_SCENE_CLASS_TYPE int32 //same as EM_SCENE_TYPE
const (
	EM_SCENE_CLASS_UNKNOW                 EM_SCENE_CLASS_TYPE = iota // unknow
	EM_SCENE_CLASS_NORMAL                                            // "Normal"
	EM_SCENE_CLASS_TRAFFIC                                           // "Traffic"
	EM_SCENE_CLASS_TRAFFIC_PATROL                                    // "TrafficPatrol"
	EM_SCENE_CLASS_FACEDETECTION                                     // "FaceDetection"
	EM_SCENE_CLASS_ATM                                               // "ATM"
	EM_SENCE_CLASS_INDOOR                                            // "Indoor"
	EM_SENCE_CLASS_FACERECOGNITION                                   // "FaceRecognition"
	EM_SENCE_CLASS_PRISON                                            // "Prison"
	EM_SENCE_CLASS_NUMBERSTAT                                        // "NumberStat"
	EM_SENCE_CLASS_HEAT_MAP                                          // "HeatMap"
	EM_SENCE_CLASS_VIDEODIAGNOSIS                                    // "VideoDiagnosis"
	EM_SENCE_CLASS_VEHICLEANALYSE                                    // "VehicleAnalyse"
	EM_SENCE_CLASS_COURSERECORD                                      // "CourseRecord"
	EM_SENCE_CLASS_VEHICLE                                           // "Vehicle"
	EM_SENCE_CLASS_STANDUPDETECTION                                  // "StandUpDetection"
	EM_SCENE_CLASS_GATE                                              // "Gate"
	EM_SCENE_CLASS_SDFACEDETECTION                                   // "SDFaceDetect"
	EM_SCENE_CLASS_HEAT_MAP_PLAN                                     // "HeatMapPlan"
	EM_SCENE_CLASS_NUMBERSTAT_PLAN                                   // "NumberStatPlan"
	EM_SCENE_CLASS_ATMFD                                             // "ATMFD"
	EM_SCENE_CLASS_HIGHWAY                                           // "Highway"
	EM_SCENE_CLASS_CITY                                              // "City"
	EM_SCENE_CLASS_LETRACK                                           // "LeTrack"
	EM_SCENE_CLASS_SCR                                               // "SCR"
	EM_SCENE_CLASS_STEREO_VISION                                     // "StereoVision"
	EM_SCENE_CLASS_HUMANDETECT                                       // "HumanDetect"
	EM_SCENE_CLASS_FACEANALYSIS                                      // "FaceAnalysis"
	EM_SCENE_CLASS_XRAY_DETECTION                                    // "XRayDetection"
	EM_SCENE_CLASS_STEREO_NUMBER                                     // "StereoNumber"
	EM_SCENE_CLASS_CROWDDISTRIMAP                                    // "CrowdDistriMap"
	EM_SCENE_CLASS_OBJECTDETECT                                      // "ObjectDetect"
	EM_SCENE_CLASS_FACEATTRIBUTE                                     // "FaceAttribute"
	EM_SCENE_CLASS_FACECOMPARE                                       // "FaceCompare"
	EM_SCENE_CLASS_STEREO_BEHAVIOR                                   // "StereoBehavior"
	EM_SCENE_CLASS_INTELLICITYMANAGER                                // "IntelliCityMgr"
	EM_SCENE_CLASS_PROTECTIVECABIN                                   // "ProtectiveCabin"
	EM_SCENE_CLASS_AIRPLANEDETECT                                    // "AirplaneDetect"
	EM_SCENE_CLASS_CROWDPOSTURE                                      // "CrowdPosture"
	EM_SCENE_CLASS_PHONECALLDETECT                                   // "PhoneCallDetect"
	EM_SCENE_CLASS_SMOKEDETECTION                                    // "SmokeDetection"
	EM_SCENE_CLASS_BOATDETECTION                                     // "BoatDetection"
	EM_SCENE_CLASS_SMOKINGDETECT                                     // "SmokingDetect"
	EM_SCENE_CLASS_WATERMONITOR                                      // "WaterMonitor"
	EM_SCENE_CLASS_GENERATEGRAPHDETECTION                            // "GenerateGraphDetection"
	EM_SCENE_CLASS_TRAFFIC_PARK                                      // "TrafficPark"
	EM_SCENE_CLASS_OPERATEMONITOR                                    // "OperateMonitor"
	EM_SCENE_CLASS_INTELLI_RETAIL                                    // IntelliRetail"
	EM_SCENE_CLASS_CLASSROOM_ANALYSE                                 // "ClassroomAnalyse"
	EM_SCENE_CLASS_FEATURE_ABSTRACT                                  // "FeatureAbstract"
	EM_SCENE_CLASS_FACEBODY_DETECT                                   // "FaceBodyDetect"
	EM_SCENE_CLASS_FACEBODY_ANALYSE                                  // "FaceBodyAnalyse"
	EM_SCENE_CLASS_VEHICLES_DISTRI                                   // "VehiclesDistri"
	EM_SCENE_CLASS_INTELLI_BREED                                     // "IntelliBreed"
	EM_SCENE_CLASS_INTELLI_PRISON                                    // "IntelliPrison"
	EM_SCENE_CLASS_ELECTRIC_DETECT                                   // "ElectricDetect"
	EM_SCENE_CLASS_RADAR_DETECT                                      // "RadarDetect"
	EM_SCENE_CLASS_PARKINGSPACE                                      // "ParkingSpace"
)

type EM_OPEN_STROBE_STATE int32

const (
	NET_OPEN_STROBE_STATE_UNKOWN EM_OPEN_STROBE_STATE = iota // 未知状态
	NET_OPEN_STROBE_STATE_CLOSE                              // 关闸
	NET_OPEN_STROBE_STATE_AUTO                               // 自动开闸
	NET_OPEN_STROBE_STATE_MANUAL                             // 手动开闸
)

type EM_VEHICLE_DIRECTION int32

const (
	NET_VEHICLE_DIRECTION_UNKOWN EM_VEHICLE_DIRECTION = iota // 未知
	NET_VEHICLE_DIRECTION_HEAD                               // 车头
	NET_VEHICLE_DIRECTION_TAIL                               // 车尾
)

type EM_NTP_STATUS int32 //NTP校时状态
const (
	NET_NTP_STATUS_UNKNOWN EM_NTP_STATUS = iota
	NET_NTP_STATUS_DISABLE
	NET_NTP_STATUS_SUCCESSFUL
	NET_NTP_STATUS_FAILED
)

type EM_COMMON_SEAT_TYPE int32

const (
	COMMON_SEAT_TYPE_UNKNOWN = 0 // 未识别
	COMMON_SEAT_TYPE_MAIN    = 1 // 主驾驶
	COMMON_SEAT_TYPE_SLAVE   = 2 // 副驾驶
)

type NET_SAFEBELT_STATE int32

const (
	SS_NUKNOW            = 0 //未知
	SS_WITH_SAFE_BELT    = 1 //已系安全带
	SS_WITHOUT_SAFE_BELT = 2 //未系安全带
)

type NET_SUNSHADE_STATE int32 //遮阳板状态
const (
	SS_NUKNOW_SUN_SHADE  NET_SUNSHADE_STATE = iota // 未知
	SS_WITH_SUN_SHADE                              // 有遮阳板
	SS_WITHOUT_SUN_SHADE                           // 无遮阳板
)

type EM_COMM_ATTACHMENT_TYPE int32 // 车辆物件类型
const (
	COMM_ATTACHMENT_TYPE_UNKNOWN    = 0 // 未知类型
	COMM_ATTACHMENT_TYPE_FURNITURE  = 1 // 摆件
	COMM_ATTACHMENT_TYPE_PENDANT    = 2 // 挂件
	COMM_ATTACHMENT_TYPE_TISSUEBOX  = 3 // 纸巾盒
	COMM_ATTACHMENT_TYPE_DANGER     = 4 // 危险品
	COMM_ATTACHMENT_TYPE_PERFUMEBOX = 5 // 香水
)

type EM_VEHICLE_TYPE int32 //收费站车型分类
const (
	EM_VEHICLE_TYPE_UNKNOWN       EM_VEHICLE_TYPE = iota // 未知
	EM_VEHICLE_TYPE_PASSENGERCAR1                        // 客1
	EM_VEHICLE_TYPE_TRUCK1                               // 货1
	EM_VEHICLE_TYPE_PASSENGERCAR2                        // 客2
	EM_VEHICLE_TYPE_TRUCK2                               // 货2
	EM_VEHICLE_TYPE_PASSENGERCAR3                        // 客3
	EM_VEHICLE_TYPE_TRUCK3                               // 货3
	EM_VEHICLE_TYPE_PASSENGERCAR4                        // 客4
	EM_VEHICLE_TYPE_TRUCK4                               // 货4
	EM_VEHICLE_TYPE_PASSENGERCAR5                        // 客5
	EM_VEHICLE_TYPE_TRUCK5                               // 货5
)

type EM_SNAPCATEGORY int32 //抓拍的类型
const (
	EM_SNAPCATEGORY_MOTOR    EM_SNAPCATEGORY = iota // 机动车
	EM_SNAPCATEGORY_NONMOTOR                        // 非机动车
)

type EM_VEHICLE_TYPE_BY_FUNC int32 // 按功能划分的车辆类型
const (
	EM_VEHICLE_TYPE_BY_FUNC_UNKNOWN              EM_VEHICLE_TYPE_BY_FUNC = iota // 未知/*以下为特种车辆类型*/
	EM_VEHICLE_TYPE_BY_FUNC_TANK_CAR                                            // 危化品车辆
	EM_VEHICLE_TYPE_BY_FUNC_SLOT_TANK_CAR                                       // 槽罐车
	EM_VEHICLE_TYPE_BY_FUNC_DREGS_CAR                                           // 渣土车
	EM_VEHICLE_TYPE_BY_FUNC_CONCRETE_MIXER_TRUCK                                // 混凝土搅拌车
	EM_VEHICLE_TYPE_BY_FUNC_TAXI                                                // 出租车
	EM_VEHICLE_TYPE_BY_FUNC_POLICE                                              // 警车
	EM_VEHICLE_TYPE_BY_FUNC_AMBULANCE                                           // 救护车
	EM_VEHICLE_TYPE_BY_FUNC_GENERAL                                             // 普通车
	EM_VEHICLE_TYPE_BY_FUNC_WATERING_CAR                                        // 洒水车
	EM_VEHICLE_TYPE_BY_FUNC_FIRE_ENGINE                                         // 消防车
	EM_VEHICLE_TYPE_BY_FUNC_MACHINESHOP_TRUCK                                   // 工程车
	EM_VEHICLE_TYPE_BY_FUNC_POWER_LOT_VEHICLE                                   // 粉粒物料车
	EM_VEHICLE_TYPE_BY_FUNC_SUCTION_SEWAGE_TRUCK                                // 吸污车
	EM_VEHICLE_TYPE_BY_FUNC_NORMAL_TANK_TRUCK                                   // 普通罐车
	EM_VEHICLE_TYPE_BY_FUNC_SCHOOL_BUS                                          // 校车
	EM_VEHICLE_TYPE_BY_FUNC_EXCAVATOR                                           // 挖掘机
	EM_VEHICLE_TYPE_BY_FUNC_BULLDOZER                                           // 推土车
	EM_VEHICLE_TYPE_BY_FUNC_CRANE                                               // 吊车
	EM_VEHICLE_TYPE_BY_FUNC_PUMP_TRUCK                                          // 泵车/*以上为特种车辆类型*/
)

type EM_STANDARD_VEHICLE_TYPE int32 // 标准车辆类型
const (
	EM_STANDARD_VEHICLE_UNKNOWN                   EM_STANDARD_VEHICLE_TYPE = iota // 未知
	EM_STANDARD_VEHICLE_MOTOR                                                     // 机动车
	EM_STANDARD_VEHICLE_BUS                                                       // 公交车
	EM_STANDARD_VEHICLE_UNLICENSED_MOTOR                                          // 无牌机动车
	EM_STANDARD_VEHICLE_LARGE_CAR                                                 // 大型汽车
	EM_STANDARD_VEHICLE_MICRO_CAR                                                 // 小型汽车
	EM_STANDARD_VEHICLE_EMBASSY_CAR                                               // 使馆汽车
	EM_STANDARD_VEHICLE_MARGINAL_CAR                                              // 领馆汽车
	EM_STANDARD_VEHICLE_AREAOUT_CAR                                               // 境外汽车
	EM_STANDARD_VEHICLE_FOREIGN_CAR                                               // 外籍汽车
	EM_STANDARD_VEHICLE_FARM_TRANS_CAR                                            // 农用运输车
	EM_STANDARD_VEHICLE_TRACTOR                                                   // 拖拉机
	EM_STANDARD_VEHICLE_TRAILER                                                   // 挂车
	EM_STANDARD_VEHICLE_COACH_CAR                                                 // 教练汽车
	EM_STANDARD_VEHICLE_TRIAL_CAR                                                 // 试验汽车
	EM_STANDARD_VEHICLE_TEMPORARYENTRY_CAR                                        // 临时入境汽车
	EM_STANDARD_VEHICLE_TEMPORARYENTRY_MOTORCYCLE                                 // 临时入境摩托
	EM_STANDARD_VEHICLE_TEMPORARY_STEER_CAR                                       // 临时行驶车
	EM_STANDARD_VEHICLE_LARGE_TRUCK                                               // 大货车
	EM_STANDARD_VEHICLE_MID_TRUCK                                                 // 中货车
	EM_STANDARD_VEHICLE_MICRO_TRUCK                                               // 小货车
	EM_STANDARD_VEHICLE_MICROBUS                                                  // 面包车
	EM_STANDARD_VEHICLE_SALOON_CAR                                                // 轿车
	EM_STANDARD_VEHICLE_CARRIAGE                                                  // 小轿车
	EM_STANDARD_VEHICLE_MINI_CARRIAGE                                             // 微型轿车
	EM_STANDARD_VEHICLE_SUV_MPV                                                   // SUV或者MPV
	EM_STANDARD_VEHICLE_SUV                                                       // SUV
	EM_STANDARD_VEHICLE_MPV                                                       // MPV
	EM_STANDARD_VEHICLE_PASSENGER_CAR                                             // 客车
	EM_STANDARD_VEHICLE_MOTOR_BUS                                                 // 大客车
	EM_STANDARD_VEHICLE_MID_PASSENGER_CAR                                         // 中客车
	EM_STANDARD_VEHICLE_MINI_BUS                                                  // 小客车
	EM_STANDARD_VEHICLE_PICKUP                                                    // 皮卡车
	EM_STANDARD_VEHICLE_OILTANK_TRUCK                                             // 油罐车
)

type EM_OVERSEA_VEHICLE_CATEGORY_TYPE int32 //海外车辆类型中的子类别，一个车辆只能是子类型的某一种。（海外专用
const (
	EM_OVERSEA_VEHICLE_CATEGORY_UNKNOWN                   EM_OVERSEA_VEHICLE_CATEGORY_TYPE = iota // 未知
	EM_OVERSEA_VEHICLE_CATEGORY_MOTORCYCLE                                                        //摩托车
	EM_OVERSEA_VEHICLE_CATEGORY_LIGHT_GOODS_VEHICLE                                               //轻型货车
	EM_OVERSEA_VEHICLE_CATEGORY_COMPANY_VEHICLE                                                   //公司用私家车
	EM_OVERSEA_VEHICLE_CATEGORY_PRIVATE_VEHICLE                                                   //个人用私家车
	EM_OVERSEA_VEHICLE_CATEGORY_TAXI                                                              //TAXI或者快线车
	EM_OVERSEA_VEHICLE_CATEGORY_TRAILER                                                           //拖车
	EM_OVERSEA_VEHICLE_CATEGORY_ENGINEERING_PLANT_VEHICLE                                         //工程车
	EM_OVERSEA_VEHICLE_CATEGORY_VERY_HEAVY_GOODS_VEHICLE                                          //超大货车
	EM_OVERSEA_VEHICLE_CATEGORY_HEAVY_GOODS_VEHICLE                                               //大货车
	EM_OVERSEA_VEHICLE_CATEGORY_PUBLIC_BUS                                                        //公共BUS
	EM_OVERSEA_VEHICLE_CATEGORY_PRIVATE_BUS                                                       //私营BUS
	EM_OVERSEA_VEHICLE_CATEGORY_SPECIAL_VEHICLE                                                   //特殊车辆
)

type EM_COLOR_TYPE int32 // 颜色类型
const (
	NET_COLOR_TYPE_RED    EM_COLOR_TYPE = iota // 红色
	NET_COLOR_TYPE_YELLOW                      // 黄色
	NET_COLOR_TYPE_GREEN                       // 绿色
	NET_COLOR_TYPE_CYAN                        // 青色
	NET_COLOR_TYPE_BLUE                        // 蓝色
	NET_COLOR_TYPE_PURPLE                      // 紫色
	NET_COLOR_TYPE_BLACK                       // 黑色
	NET_COLOR_TYPE_WHITE                       // 白色
	NET_COLOR_TYPE_MAX
)

type EM_MSG_OBJ_PERSON_DIRECTION int32 // 入侵方向
const (
	EM_MSG_OBJ_PERSON_DIRECTION_UNKOWN        EM_MSG_OBJ_PERSON_DIRECTION = iota // 未知方向
	EM_MSG_OBJ_PERSON_DIRECTION_LEFT_TO_RIGHT                                    // 从左向右
	EM_MSG_OBJ_PERSON_DIRECTION_RIGHT_TO_LEFT                                    // 从右向左
)

type EM_CAR_TYPE int32 // 车辆类型
const (
	EM_CAR_0   EM_CAR_TYPE = iota //其他车辆
	EM_CAR_1                      //大型普通客车
	EM_CAR_2                      //大型双层客车
	EM_CAR_3                      //大型卧铺客车
	EM_CAR_4                      //大型铰接客车
	EM_CAR_5                      //大型越野客车
	EM_CAR_6                      //大型轿车
	EM_CAR_7                      //大型专用客车
	EM_CAR_8                      //大型专用校车
	EM_CAR_9                      //中型普通客车
	EM_CAR_10                     //中型双层客车
	EM_CAR_11                     //中型卧铺客车
	EM_CAR_12                     //中型铰接客车
	EM_CAR_13                     //中型越野客车
	EM_CAR_14                     //中型轿车
	EM_CAR_15                     //中型专用客车
	EM_CAR_16                     //中型专用校车
	EM_CAR_17                     //小型普通客车
	EM_CAR_18                     //小型越野客车
	EM_CAR_19                     //小型轿车
	EM_CAR_20                     //小型专用客车
	EM_CAR_21                     //小型专用校车
	EM_CAR_22                     //小型面包车
	EM_CAR_23                     //微型普通客车
	EM_CAR_24                     //微型越野客车
	EM_CAR_25                     //微型轿车
	EM_CAR_26                     //微型面包车
	EM_CAR_27                     //重型半挂牵引车
	EM_CAR_28                     //重型全挂牵引车
	EM_CAR_29                     //中型半挂牵引车
	EM_CAR_30                     //中型全挂牵引车
	EM_CAR_31                     //轻型半挂牵引车
	EM_CAR_32                     //轻型全挂牵引车
	EM_CAR_33                     //大型非载货专项作业车
	EM_CAR_34                     //大型载货专项作业车
	EM_CAR_35                     //中型非载货专项作业车
	EM_CAR_36                     //中型载货专项作业车
	EM_CAR_37                     //小型非载货专项作业车
	EM_CAR_38                     //小型载货专项作业车
	EM_CAR_39                     //微型非载货专项作业车
	EM_CAR_40                     //微型载货专项作业车
	EM_CAR_41                     //重型非载货专项作业车
	EM_CAR_42                     //重型载货专项作业车
	EM_CAR_43                     //轻型非载货专项作业车
	EM_CAR_44                     //轻型载货专项作业车
	EM_CAR_45                     //普通正三轮摩托车
	EM_CAR_46                     //轻便正三轮摩托车
	EM_CAR_47                     //正三轮载客摩托车
	EM_CAR_48                     //正三轮载货摩托车
	EM_CAR_49                     //侧三轮摩托车
	EM_CAR_50                     //普通二轮摩托车
	EM_CAR_51                     //轻便二轮摩托车
	EM_CAR_52                     //无轨电车
	EM_CAR_53                     //有轨电车
	EM_CAR_54                     //三轮汽车
	EM_CAR_55                     //轮式装载机械
	EM_CAR_56                     //轮式挖掘机械
	EM_CAR_57                     //轮式平地机械
	EM_CAR_58                     //重型普通货车
	EM_CAR_59                     //重型厢式货车
	EM_CAR_60                     //重型封闭货车
	EM_CAR_61                     //重型罐式货车
	EM_CAR_62                     //重型平板货车
	EM_CAR_63                     //重型集装箱车
	EM_CAR_64                     //重型自卸货车
	EM_CAR_65                     //重型特殊结构货车
	EM_CAR_66                     //重型仓栅式货车
	EM_CAR_67                     //重型车辆运输车
	EM_CAR_68                     //重型厢式自卸货车
	EM_CAR_69                     //重型罐式自卸货车
	EM_CAR_70                     //重型平板自卸货车
	EM_CAR_71                     //重型集装箱自卸货车
	EM_CAR_72                     //重型特殊结构自卸货车
	EM_CAR_73                     //重型仓栅式自卸货车
	EM_CAR_74                     //中型普通货车
	EM_CAR_75                     //中型厢式货车
	EM_CAR_76                     //中型封闭货车
	EM_CAR_77                     //中型罐式货车
	EM_CAR_78                     //中型平板货车
	EM_CAR_79                     //中型集装箱车
	EM_CAR_80                     //中型自卸货车
	EM_CAR_81                     //中型特殊结构货车
	EM_CAR_82                     //中型仓栅式货车
	EM_CAR_83                     //中型车辆运输车
	EM_CAR_84                     //中型厢式自卸货车
	EM_CAR_85                     //中型罐式自卸货车
	EM_CAR_86                     //中型平板自卸货车
	EM_CAR_87                     //中型集装箱自卸货车
	EM_CAR_88                     //中型特殊结构自卸货车
	EM_CAR_89                     //中型仓栅式自卸货车
	EM_CAR_90                     //轻型普通货车
	EM_CAR_91                     //轻型厢式货车
	EM_CAR_92                     //轻型封闭货车
	EM_CAR_93                     //轻型罐式货车
	EM_CAR_94                     //轻型平板货车
	EM_CAR_95                     //轻型自卸货车
	EM_CAR_96                     //轻型特殊结构货车
	EM_CAR_97                     //轻型仓栅式货车
	EM_CAR_98                     //轻型车辆运输车
	EM_CAR_99                     //轻型厢式自卸货车
	EM_CAR_100                    //轻型罐式自卸货车
	EM_CAR_101                    //轻型平板自卸货车
	EM_CAR_102                    //轻型特殊结构自卸货车
	EM_CAR_103                    //轻型仓栅式自卸货车
	EM_CAR_104                    //微型普通货车
	EM_CAR_105                    //微型厢式货车
	EM_CAR_106                    //微型封闭货车
	EM_CAR_107                    //微型罐式货车
	EM_CAR_108                    //微型自卸货车
	EM_CAR_109                    //微型特殊结构货车
	EM_CAR_110                    //微型仓栅式货车
	EM_CAR_111                    //微型车辆运输车
	EM_CAR_112                    //微型厢式自卸货车
	EM_CAR_113                    //微型罐式自卸货车
	EM_CAR_114                    //微型特殊结构自卸货车
	EM_CAR_115                    //微型仓栅式自卸货车
	EM_CAR_116                    //普通低速货车
	EM_CAR_117                    //厢式低速货车
	EM_CAR_118                    //罐式低速货车
	EM_CAR_119                    //自卸低速货车
	EM_CAR_120                    //仓栅式低速货车
	EM_CAR_121                    //厢式自卸低速货车
	EM_CAR_122                    //罐式自卸低速货车
	EM_CAR_123                    //重型普通全挂车
	EM_CAR_124                    //重型厢式全挂车
	EM_CAR_125                    //重型罐式全挂车
	EM_CAR_126                    //重型平板全挂车
	EM_CAR_127                    //重型集装箱全挂车
	EM_CAR_128                    //重型自卸全挂车
	EM_CAR_129                    //重型仓栅式全挂车
	EM_CAR_130                    //重型旅居全挂车
	EM_CAR_131                    //重型专项作业全挂车
	EM_CAR_132                    //重型厢式自卸全挂车
	EM_CAR_133                    //重型罐式自卸全挂车
	EM_CAR_134                    //重型平板自卸全挂车
	EM_CAR_135                    //重型集装箱自卸全挂车
	EM_CAR_136                    //重型仓栅式自卸全挂车
	EM_CAR_137                    //重型专项作业自卸全挂车
	EM_CAR_138                    //中型普通全挂车
	EM_CAR_139                    //中型厢式全挂车
	EM_CAR_140                    //中型罐式全挂车
	EM_CAR_141                    //中型平板全挂车
	EM_CAR_142                    //中型集装箱全挂车
	EM_CAR_143                    //中型自卸全挂车
	EM_CAR_144                    //中型仓栅式全挂车
	EM_CAR_145                    //中型旅居全挂车
	EM_CAR_146                    //中型专项作业全挂车
	EM_CAR_147                    //中型厢式自卸全挂车
	EM_CAR_148                    //中型罐式自卸全挂车
	EM_CAR_149                    //中型平板自卸全挂车
	EM_CAR_150                    //中型集装箱自卸全挂车
	EM_CAR_151                    //中型仓栅式自卸全挂车
	EM_CAR_152                    //中型专项作业自卸全挂车
	EM_CAR_153                    //轻型普通全挂车
	EM_CAR_154                    //轻型厢式全挂车
	EM_CAR_155                    //轻型罐式全挂车
	EM_CAR_156                    //轻型平板全挂车
	EM_CAR_157                    //轻型自卸全挂车
	EM_CAR_158                    //轻型仓栅式全挂车
	EM_CAR_159                    //轻型旅居全挂车
	EM_CAR_160                    //轻型专项作业全挂车
	EM_CAR_161                    //轻型厢式自卸全挂车
	EM_CAR_162                    //轻型罐式自卸全挂车
	EM_CAR_163                    //轻型平板自卸全挂车
	EM_CAR_164                    //轻型集装箱自卸全挂车
	EM_CAR_165                    //轻型仓栅式自卸全挂车
	EM_CAR_166                    //轻型专项作业自卸全挂车
	EM_CAR_167                    //重型普通半挂车
	EM_CAR_168                    //重型厢式半挂车
	EM_CAR_169                    //重型罐式半挂车
	EM_CAR_170                    //重型平板半挂车
	EM_CAR_171                    //重型集装箱半挂车
	EM_CAR_172                    //重型自卸半挂车
	EM_CAR_173                    //重型特殊结构半挂车
	EM_CAR_174                    //重型仓栅式半挂车
	EM_CAR_175                    //重型旅居半挂车
	EM_CAR_176                    //重型专项作业半挂车
	EM_CAR_177                    //重型低平板半挂车
	EM_CAR_178                    //重型车辆运输半挂车
	EM_CAR_179                    //重型罐式自卸半挂车
	EM_CAR_180                    //重型平板自卸半挂车
	EM_CAR_181                    //重型集装箱自卸半挂车
	EM_CAR_182                    //重型特殊结构自卸半挂车
	EM_CAR_183                    //重型仓栅式自卸半挂车
	EM_CAR_184                    //重型专项作业自卸半挂车
	EM_CAR_185                    //重型低平板自卸半挂车
	EM_CAR_186                    //重型中置轴旅居挂车
	EM_CAR_187                    //重型中置轴车辆运输车
	EM_CAR_188                    //重型中置轴普通挂车
	EM_CAR_189                    //中型普通半挂车
	EM_CAR_190                    //中型厢式半挂车
	EM_CAR_191                    //中型罐式半挂车
	EM_CAR_192                    //中型平板半挂车
	EM_CAR_193                    //中型集装箱半挂车
	EM_CAR_194                    //中型自卸半挂车
	EM_CAR_195                    //中型特殊结构半挂车
	EM_CAR_196                    //中型仓栅式半挂车
	EM_CAR_197                    //中型旅居半挂车
	EM_CAR_198                    //中型专项作业半挂车
	EM_CAR_199                    //中型低平板半挂车
	EM_CAR_200                    //中型车辆运输半挂车
	EM_CAR_201                    //中型罐式自卸半挂车
	EM_CAR_202                    //中型平板自卸半挂车
	EM_CAR_203                    //中型集装箱自卸半挂车
	EM_CAR_204                    //中型特殊结构自卸挂车
	EM_CAR_205                    //中型仓栅式自卸半挂车
	EM_CAR_206                    //中型专项作业自卸半挂车
	EM_CAR_207                    //中型低平板自卸半挂车
	EM_CAR_208                    //中型中置轴旅居挂车
	EM_CAR_209                    //中型中置轴车辆运输车
	EM_CAR_210                    //中型中置轴普通挂车
	EM_CAR_211                    //轻型普通半挂车
	EM_CAR_212                    //轻型厢式半挂车
	EM_CAR_213                    //轻型罐式半挂车
	EM_CAR_214                    //轻型平板半挂车
	EM_CAR_215                    //轻型自卸半挂车
	EM_CAR_216                    //轻型仓栅式半挂车
	EM_CAR_217                    //轻型旅居半挂车
	EM_CAR_218                    //轻型专项作业半挂车
	EM_CAR_219                    //轻型低平板半挂车
	EM_CAR_220                    //轻型车辆运输半挂车
	EM_CAR_221                    //轻型罐式自卸半挂车
	EM_CAR_222                    //轻型平板自卸半挂车
	EM_CAR_223                    //轻型集装箱自卸半挂车
	EM_CAR_224                    //轻型特殊结构自卸挂车
	EM_CAR_225                    //轻型仓栅式自卸半挂车
	EM_CAR_226                    //轻型专项作业自卸半挂车
	EM_CAR_227                    //轻型低平板自卸半挂车
	EM_CAR_228                    //轻型中置轴旅居挂车
	EM_CAR_229                    //轻型中置轴车辆运输车
	EM_CAR_230                    //轻型中置轴普通挂车
)

type EM_PLATE_TYPE int32 // 号牌类型
const (
	EM_PLATE_OTHER                                        EM_PLATE_TYPE = iota // 其他车
	EM_PLATE_BIG_CAR                                                           // 大型汽车
	EM_PLATE_SMALL_CAR                                                         // 小型汽车
	EM_PLATE_EMBASSY_CAR                                                       // 使馆汽车
	EM_PLATE_CONSULATE_CAR                                                     // 领馆汽车
	EM_PLATE_ABROAD_CAR                                                        // 境外汽车
	EM_PLATE_FOREIGN_CAR                                                       // 外籍汽车
	EM_PLATE_LOW_SPEED_CAR                                                     // 低速车
	EM_PLATE_COACH_CAR                                                         // 教练车
	EM_PLATE_MOTORCYCLE                                                        // 摩托车
	EM_PLATE_NEW_POWER_CAR                                                     // 新能源车
	EM_PLATE_POLICE_CAR                                                        // 警用车
	EM_PLATE_HONGKONG_MACAO_CAR                                                // 港澳两地车
	EM_PLATE_ARMEDPOLICE_CAR                                                   // 武警车辆
	EM_PLATE_MILITARY_CAR                                                      // 军队车辆
	EM_PLATE_TEMPORARY_LICENSE_FOR_NON_MOTOR_VEHICLES                          // 非机动车临时牌照
	EM_PLATE_OFFICIAL_LICENSE_PLATE_FOR_NON_MOTOR_VEHICLE                      // 非机动车牌正式牌照
)

type EM_CAR_COLOR_TYPE int32 // 车身颜色
const (
	EM_CAR_COLOR_WHITE        EM_CAR_COLOR_TYPE = iota // 白色
	EM_CAR_COLOR_BLACK                                 // 黑色
	EM_CAR_COLOR_RED                                   // 红色
	EM_CAR_COLOR_YELLOW                                // 黄色
	EM_CAR_COLOR_GRAY                                  // 灰色
	EM_CAR_COLOR_BLUE                                  // 蓝色
	EM_CAR_COLOR_GREEN                                 // 绿色
	EM_CAR_COLOR_PINK                                  // 粉色
	EM_CAR_COLOR_PURPLE                                // 紫色
	EM_CAR_COLOR_DARK_PURPLE                           // 暗紫色
	EM_CAR_COLOR_BROWN                                 // 棕色
	EM_CAR_COLOR_MAROON                                // 粟色
	EM_CAR_COLOR_SILVER_GRAY                           // 银灰色
	EM_CAR_COLOR_DARK_GRAY                             // 暗灰色
	EM_CAR_COLOR_WHITE_SMOKE                           // 白烟色
	EM_CAR_COLOR_DEEP_ORANGE                           // 深橙色
	EM_CAR_COLOR_LIGHT_ROSE                            // 浅玫瑰色
	EM_CAR_COLOR_TOMATO_RED                            // 番茄红色
	EM_CAR_COLOR_OLIVE                                 // 橄榄色
	EM_CAR_COLOR_GOLDEN                                // 金色
	EM_CAR_COLOR_DARK_OLIVE                            // 暗橄榄色
	EM_CAR_COLOR_YELLOW_GREEN                          // 黄绿色
	EM_CAR_COLOR_GREEN_YELLOW                          // 绿黄色
	EM_CAR_COLOR_FOREST_GREEN                          // 森林绿
	EM_CAR_COLOR_OCEAN_BLUE                            // 海洋绿
	EM_CAR_COLOR_DEEP_SKYBLUE                          // 深天蓝
	EM_CAR_COLOR_CYAN                                  // 青色
	EM_CAR_COLOR_DEEP_BLUE                             // 深蓝色
	EM_CAR_COLOR_DEEP_RED                              // 深红色
	EM_CAR_COLOR_DEEP_GREEN                            // 深绿色
	EM_CAR_COLOR_DEEP_YELLOW                           // 深黄色
	EM_CAR_COLOR_DEEP_PINK                             // 深粉色
	EM_CAR_COLOR_DEEP_PURPLE                           // 深紫色
	EM_CAR_COLOR_DEEP_BROWN                            // 深棕色
	EM_CAR_COLOR_DEEP_CYAN                             // 深青色
	EM_CAR_COLOR_ORANGE                                // 橙色
	EM_CAR_COLOR_DEEP_GOLDEN                           // 深金色
	EM_CAR_COLOR_OTHER        = 255                    // 未识别、其他
)

type EM_USE_PROPERTY_TYPE int32 // 使用性质
const (
	EM_USE_PROPERTY_NONOPERATING                   EM_USE_PROPERTY_TYPE = iota // 非营运
	EM_USE_PROPERTY_HIGWAY                                                     // 公路客运,旅游客运
	EM_USE_PROPERTY_BUS                                                        // 公交客运
	EM_USE_PROPERTY_TAXI                                                       // 出租客运
	EM_USE_PROPERTY_FREIGHT                                                    // 货运
	EM_USE_PROPERTY_LEASE                                                      // 租赁
	EM_USE_PROPERTY_SECURITY                                                   // 警用,消防,救护,工程救险
	EM_USE_PROPERTY_COACH                                                      // 教练
	EM_USE_PROPERTY_SCHOOLBUS                                                  // 幼儿校车,小学生校车,其他校车
	EM_USE_PROPERTY_FOR_DANGE_VEHICLE                                          // 危化品运输
	EM_USE_PROPERTY_OTHER                                                      // 其他
	EM_USE_PROPERTY_ONLINE_CAR_HAILING                                         // 网约车
	EM_USE_PROPERTY_NON_MOTORIZED_TAKE_OUT_VEHICLE                             // 非机动外卖车
	EM_USE_PROPERTY_NON_MOTORIZED_EXPRESS_CAR                                  // 非机动快递车
)

type EM_CARD_PROVINCE int32 // 卡号省份
const (
	EM_CARD_UNKNOWN        = 10 // 解析出错，未知省份
	EM_CARD_BEIJING        = 11 // 北京
	EM_CARD_TIANJIN        = 12 // 天津
	EM_CARD_HEBEI          = 13 // 河北
	EM_CARD_SHANXI_TAIYUAN = 14 // 山西
	EM_CARD_NEIMENGGU      = 15 // 内蒙古
	EM_CARD_LIAONING       = 21 // 辽宁
	EM_CARD_JILIN          = 22 // 吉林
	EM_CARD_HEILONGJIANG   = 23 // 黑龙江
	EM_CARD_SHANGHAI       = 31 // 上海
	EM_CARD_JIANGSU        = 32 // 江苏
	EM_CARD_ZHEJIANG       = 33 // 浙江
	EM_CARD_ANHUI          = 34 // 安徽
	EM_CARD_FUJIAN         = 35 // 福建
	EM_CARD_JIANGXI        = 36 // 江西
	EM_CARD_SHANDONG       = 37 // 山东
	EM_CARD_HENAN          = 41 // 河南
	EM_CARD_HUBEI          = 42 // 湖北
	EM_CARD_HUNAN          = 43 // 湖南
	EM_CARD_GUANGDONG      = 44 // 广东
	EM_CARD_GUANGXI        = 45 // 广西
	EM_CARD_HAINAN         = 46 // 海南
	EM_CARD_CHONGQING      = 50 // 重庆
	EM_CARD_SICHUAN        = 51 // 四川
	EM_CARD_GUIZHOU        = 52 // 贵州
	EM_CARD_YUNNAN         = 53 // 云南
	EM_CARD_XIZANG         = 54 // 西藏
	EM_CARD_SHANXI_XIAN    = 61 // 陕西
	EM_CARD_GANSU          = 62 // 甘肃
	EM_CARD_QINGHAI        = 63 // 青海
	EM_CARD_NINGXIA        = 64 // 宁夏
	EM_CARD_XINJIANG       = 65 // 新疆
	EM_CARD_XIANGGANG      = 71 // 香港
	EM_CARD_AOMEN          = 82 // 澳门
)

type EM_CLOTHES_COLOR int32 // 衣服颜色
const (
	EM_CLOTHES_COLOR_UNKNOWN EM_CLOTHES_COLOR = iota // 未知
	EM_CLOTHES_COLOR_WHITE                           // 白色
	EM_CLOTHES_COLOR_ORANGE                          // 橙色
	EM_CLOTHES_COLOR_PINK                            // 粉色
	EM_CLOTHES_COLOR_BLACK                           // 黑色
	EM_CLOTHES_COLOR_RED                             // 红色
	EM_CLOTHES_COLOR_YELLOW                          // 黄色
	EM_CLOTHES_COLOR_GRAY                            // 灰色
	EM_CLOTHES_COLOR_BLUE                            // 蓝色
	EM_CLOTHES_COLOR_GREEN                           // 绿色
	EM_CLOTHES_COLOR_PURPLE                          // 紫色
	EM_CLOTHES_COLOR_BROWN                           // 棕色
	EM_CLOTHES_COLOR_OTHER                           // 其他颜色
)

type EM_COAT_TYPE int32 //上衣类型
const (
	EM_COAT_TYPE_UNKNOWN     EM_COAT_TYPE = iota // 未知
	EM_COAT_TYPE_LONG_SLEEVE                     // 长袖
	EM_COAT_TYPE_COTTA                           // 短袖
)

type EM_TROUSERS_TYPE int32 // 裤子类型
const (
	EM_TROUSERS_TYPE_UNKNOWN  EM_TROUSERS_TYPE = iota // 未知
	EM_TROUSERS_TYPE_TROUSERS                         // 长裤
	EM_TROUSERS_TYPE_SHORTS                           // 短裤
	EM_TROUSERS_TYPE_SKIRT                            // 裙子
)

type EM_HAS_HAT int32 // 是否戴帽子
const (
	EM_HAS_HAT_UNKNOWN EM_HAS_HAT = iota // 未知
	EM_HAS_HAT_NO                        // 不戴帽子
	EM_HAS_HAT_YES                       // 戴帽子
)

type EM_HAS_BAG int32 // 是否戴包(包括背包或拎包)
const (
	EM_HAS_BAG_UNKNOWN EM_HAS_BAG = iota // 未知
	EM_HAS_BAG_NO                        // 不带包
	EM_HAS_BAG_YES                       // 带包
)

type EM_HAS_BACK_BAG int32 // 是否有背包
const (
	EM_HAS_BACK_BAG_UNKNOWN EM_HAS_BACK_BAG = iota // 未知
	EM_HAS_BACK_BAG_NO                             // 没有背包
	EM_HAS_BACK_BAG_YES                            // 有背包
)

type EM_HAS_CARRIER_BAG int32 // 是否有手提包
const (
	EM_HAS_CARRIER_BAG_UNKNOWN EM_HAS_CARRIER_BAG = iota // 未知
	EM_HAS_CARRIER_BAG_NO                                // 没有手提包
	EM_HAS_CARRIER_BAG_YES                               // 有手提包
)

type EM_HAS_SHOULDER_BAG int32 // 是否有肩包
const (
	EM_HAS_SHOULDER_BAG_UNKNOWN EM_HAS_SHOULDER_BAG = iota // 未知
	EM_HAS_SHOULDER_BAG_NO                                 // 没有肩包
	EM_HAS_SHOULDER_BAG_YES                                // 有肩包
)

type EM_HAS_MESSENGER_BAG int32 // 是否有斜挎包
const (
	EM_HAS_MESSENGER_BAG_UNKNOWN EM_HAS_MESSENGER_BAG = iota // 未知
	EM_HAS_MESSENGER_BAG_NO                                  // 没有斜挎包
	EM_HAS_MESSENGER_BAG_YES                                 // 有斜挎包
)

type EM_ANGLE_TYPE int32 // 角度
const (
	EM_ANGLE_UNKNOWN EM_ANGLE_TYPE = iota // 未知
	EM_ANGLE_FRONT                        // 正面
	EM_ANGLE_SIDE                         // 侧面
	EM_ANGLE_BACK                         // 背面
)

type EM_HAS_UMBRELLA int32 // 是否打伞
const (
	EM_HAS_UMBRELLA_UNKNOWN EM_HAS_UMBRELLA = iota // 未知
	EM_HAS_UMBRELLA_NO                             // 未打伞
	EM_HAS_UMBRELLA_YES                            // 打伞
)

type EM_BAG_TYPE int32 // 包类型
const (
	EM_BAG_UNKNOWN     EM_BAG_TYPE = iota // 未知
	EM_BAG_HANDBAG                        // 手提包
	EM_BAG_SHOULDERBAG                    // 肩包
	EM_BAG_KNAPSACK                       // 背包
	EM_BAG_DRAWBARBOX                     // 拉杆箱
	EM_BAG_WAISTPACK                      // 腰包
	EM_BAG_NONE                           // 无包
)

type EM_CAP_TYPE int32 // 帽子类型
const (
	EM_CAP_UNKNOWN  EM_CAP_TYPE = iota // 未知
	EM_CAP_ORDINARY                    // 普通帽子
	EM_CAP_HELMET                      // 头盔
	EM_CAP_SAFE                        // 安全帽
)

type EM_CLOTHES_PATTERN int32 // 衣服图案
const (
	EM_CLOTHES_PATTERN_UNKNOWN   EM_CLOTHES_PATTERN = iota // 未知
	EM_CLOTHES_PATTERN_PURE                                // 纯色
	EM_CLOTHES_PATTERN_STRIPE                              // 条纹
	EM_CLOTHES_PATTERN_PATTERN                             // 图案
	EM_CLOTHES_PATTERN_GAP                                 // 缝隙
	EM_CLOTHES_PATTERN_LATTICE                             // 格子
	EM_CLOTHES_PATTERN_SPLITJOIN                           // 拼接
)

type EM_HAIR_STYLE int32 // 头发样式
const (
	EM_HAIR_UNKNOWN      EM_HAIR_STYLE = iota // 未知
	EM_HAIR_LONG_HAIR                         // 长发
	EM_HAIR_SHORT_HAIR                        // 短发
	EM_HAIR_PONYTAIL                          // 马尾
	EM_HAIR_UPDO                              // 盘发
	EM_HAIR_HEAD_BLOCKED                      // 头部被遮挡
	EM_HAIR_NONE                              // 无头发
)

type EM_SEX_TYPE int32 // 性别
const (
	EM_SEX_TYPE_UNKNOWN EM_SEX_TYPE = iota //未知
	EM_SEX_TYPE_MALE                       //男性
	EM_SEX_TYPE_FEMALE                     //女性
)

type EM_GLASSES_TYPE int32 // 眼镜类型
const (
	EM_GLASSES_UNKNOWN  EM_GLASSES_TYPE = iota // 未知
	EM_GLASSES_SUNGLASS                        // 太阳眼镜
	EM_GLASSES_GLASS                           // 普通眼镜
)

type EM_RACE_TYPE int32 // 种族类型
const (
	EM_RACE_UNKNOWN EM_RACE_TYPE = iota // 未知
	EM_RACE_NODISTI                     // 未识别
	EM_RACE_YELLOW                      // 黄种人
	EM_RACE_BLACK                       // 黑人
	EM_RACE_WHITE                       // 白人
)

type EM_EYE_STATE_TYPE int32 // 眼睛状态
const (
	EM_EYE_STATE_UNKNOWN EM_EYE_STATE_TYPE = iota // 未知
	EM_EYE_STATE_NODISTI                          // 未识别
	EM_EYE_STATE_CLOSE                            // 闭眼
	EM_EYE_STATE_OPEN                             // 睁眼
)

type EM_MOUTH_STATE_TYPE int32 // 嘴巴状态
const (
	EM_MOUTH_STATE_UNKNOWN EM_MOUTH_STATE_TYPE = iota // 未知
	EM_MOUTH_STATE_NODISTI                            // 未识别
	EM_MOUTH_STATE_CLOSE                              // 闭嘴
	EM_MOUTH_STATE_OPEN                               // 张嘴
)

type EM_MASK_STATE_TYPE int32 // 口罩状态
const (
	EM_MASK_STATE_UNKNOWN EM_MASK_STATE_TYPE = iota // 未知
	EM_MASK_STATE_NODISTI                           // 未识别
	EM_MASK_STATE_NOMASK                            // 没戴口罩
	EM_MASK_STATE_WEAR                              // 戴口罩
)

type EM_BEARD_STATE_TYPE int32 // 胡子状态
const (
	EM_BEARD_STATE_UNKNOWN   EM_BEARD_STATE_TYPE = iota // 未知
	EM_BEARD_STATE_NODISTI                              // 未识别
	EM_BEARD_STATE_NOBEARD                              // 没胡子
	EM_BEARD_STATE_HAVEBEARD                            // 有胡子
)

type EM_FACEDETECT_GLASSES_TYPE int32 // 人脸检测人脸戴眼镜特征类型
const (
	EM_FACEDETECT_GLASSES_UNKNOWN EM_FACEDETECT_GLASSES_TYPE = iota // 未知
	EM_FACEDETECT_WITH_GLASSES                                      // 戴眼镜
	EM_FACEDETECT_WITHOUT_GLASSES                                   // 不戴眼镜
)

type EM_DEV_EVENT_FACEDETECT_SEX_TYPE int32 //人脸检测对应性别类型
const (
	EM_DEV_EVENT_FACEDETECT_SEX_TYPE_UNKNOWN EM_DEV_EVENT_FACEDETECT_SEX_TYPE = iota // 未知
	EM_DEV_EVENT_FACEDETECT_SEX_TYPE_MAN                                             // 男性
	EM_DEV_EVENT_FACEDETECT_SEX_TYPE_WOMAN                                           // 女性
)

type EM_NATION_TYPE int32 //民族
const (
	EM_NATION_TYPE_UNKNOWN      EM_NATION_TYPE = iota // 未知
	EM_NATION_TYPE_UYGUR                              // 维族(新疆)
	EM_NATION_TYPE_OTHER                              // 其他
	EM_NATION_TYPE_UNIDENTIFIED                       // 设备未识别
)

type EM_CLASS_TYPE int32 // 大类业务方案
const (
	EM_CLASS_UNKNOWN            = 0  // 未知业务
	EM_CLASS_VIDEO_SYNOPSIS     = 1  // 视频浓缩
	EM_CLASS_TRAFFIV_GATE       = 2  //卡口
	EM_CLASS_ELECTRONIC_POLICE  = 3  // 电警
	EM_CLASS_SINGLE_PTZ_PARKING = 4  // 单球违停
	EM_CLASS_PTZ_PARKINBG       = 5  // 主从违停
	EM_CLASS_TRAFFIC            = 6  // 交通事件"Traffic"
	EM_CLASS_NORMAL             = 7  // 通用行为分析"Normal"
	EM_CLASS_PRISON             = 8  // 监所行为分析"Prison"
	EM_CLASS_ATM                = 9  // 金融行为分析"ATM"
	EM_CLASS_METRO              = 10 // 地铁行为分析
	EM_CLASS_FACE_DETECTION     = 11 // 人脸检测"FaceDetection"
	EM_CLASS_FACE_RECOGNITION   = 12 // 人脸识别"FaceRecognition"
	EM_CLASS_NUMBER_STAT        = 13 // 人数统计"NumberStat"
	EM_CLASS_HEAT_MAP           = 14 // 热度图"HeatMap"
	EM_CLASS_VIDEO_DIAGNOSIS    = 15 // 视频诊断"VideoDiagnosis"
	EM_CLASS_VIDEO_ENHANCE      = 16 // 视频增强
	EM_CLASS_SMOKEFIRE_DETECT   = 17 // 烟火检测
	EM_CLASS_VEHICLE_ANALYSE    = 18 // 车辆特征识别"VehicleAnalyse"
	EM_CLASS_PERSON_FEATURE     = 19 // 人员特征识别
	EM_CLASS_SDFACEDETECTION    = 20 // 多预置点人脸检测"SDFaceDetect"
	//配置一条规则但可以在不同预置点下生效
	EM_CLASS_HEAT_MAP_PLAN          = 21 // 球机热度图计划"HeatMapPlan"
	EM_CLASS_NUMBERSTAT_PLAN        = 22 // 球机客流量统计计划 "NumberStatPlan"
	EM_CLASS_ATMFD                  = 23 // 金融人脸检测，包括正常人脸、异常人脸、相邻人脸、头盔人脸等针对ATM场景特殊优化
	EM_CLASS_HIGHWAY                = 24 // 高速交通事件检测"Highway"
	EM_CLASS_CITY                   = 25 // 城市交通事件检测 "City"
	EM_CLASS_LETRACK                = 26 // 民用简易跟踪"LeTrack"
	EM_CLASS_SCR                    = 27 // 打靶相机"SCR"
	EM_CLASS_STEREO_VISION          = 28 // 立体视觉(双目)"StereoVision"
	EM_CLASS_HUMANDETECT            = 29 // 人体检测"HumanDetect"
	EM_CLASS_FACE_ANALYSIS          = 30 // 人脸分析 "FaceAnalysis"
	EM_CALSS_XRAY_DETECTION         = 31 // X光检测 "XRayDetection"
	EM_CLASS_STEREO_NUMBER          = 32 // 双目相机客流量统计 "StereoNumber"
	EM_CLASS_CROWDDISTRIMAP         = 33 // 人群分布图
	EM_CLASS_OBJECTDETECT           = 34 // 目标检测
	EM_CLASS_FACEATTRIBUTE          = 35 // IVSS人脸检测 "FaceAttribute"
	EM_CLASS_FACECOMPARE            = 36 // IVSS人脸识别 "FaceCompare"
	EM_CALSS_STEREO_BEHAVIOR        = 37 // 立体行为分析 "StereoBehavior"
	EM_CALSS_INTELLICITYMANAGER     = 38 // 智慧城管 "IntelliCityMgr"
	EM_CALSS_PROTECTIVECABIN        = 39 // 防护舱（ATM舱内）"ProtectiveCabin"
	EM_CALSS_AIRPLANEDETECT         = 40 // 飞机行为检测 "AirplaneDetect"
	EM_CALSS_CROWDPOSTURE           = 41 // 人群态势（人群分布图服务）"CrowdPosture"
	EM_CLASS_PHONECALLDETECT        = 42 // 打电话检测 "PhoneCallDetect"
	EM_CLASS_SMOKEDETECTION         = 43 // 烟雾检测 "SmokeDetection"
	EM_CLASS_BOATDETECTION          = 44 // 船只检测 "BoatDetection"
	EM_CLASS_SMOKINGDETECT          = 45 // 吸烟检测 "SmokingDetect"
	EM_CLASS_WATERMONITOR           = 46 // 水利监测 "WaterMonitor"
	EM_CLASS_GENERATEGRAPHDETECTION = 47 // 生成图规则 "GenerateGraphDetection"
	EM_CLASS_TRAFFIC_PARK           = 48 // 交通停车 "TrafficPark"
	EM_CLASS_OPERATEMONITOR         = 49 // 作业检测 "OperateMonitor"
	EM_CLASS_INTELLI_RETAIL         = 50 // 智慧零售大类 "IntelliRetail"
	EM_CLASS_CLASSROOM_ANALYSE      = 51 // 教育智慧课堂"ClassroomAnalyse"
	EM_CLASS_FEATURE_ABSTRACT       = 52 // 特征向量提取大类 "FeatureAbstract"
	EM_CLASS_FACEBODY_DETECT        = 53 // 人体检测大类 "FaceBodyDetect"
	EM_CLASS_FACEBODY_ANALYSE       = 54 // 人体识别大类 "FaceBodyAnalyse"
	EM_CLASS_VEHICLES_DISTRI        = 55 // 车辆密度 "VehiclesDistri"
	EM_CLASS_INTELLI_BREED          = 56 // 智慧养殖检测 "IntelliBreed"
	EM_CLASS_INTELLI_PRISON         = 57 // 监狱行为分析 "IntelliPrison"
	EM_CLASS_ELECTRIC_DETECT        = 58 // 电力检测 "ElectricDetect"
	EM_CLASS_RADAR_DETECT           = 59 // 雷达检测 "RadarDetect"
	EM_CLASS_PARKINGSPACE           = 60 // 车位检测大类 "ParkingSpace"
	EM_CLASS_INTELLI_FINANCE        = 61 // 智慧金融 "IntelliFinance"
	EM_CLASS_CROWD_ABNORMAL         = 62 // 人群异常检测 "CrowdAbnormal"
	EM_CLASS_ANATOMY_TEMP_DETECT    = 63 // 人体温智能检测 "AnatomyTempDetect"
)

// LoginWithHighLevelSecurity input param
type NET_IN_LOGIN_WITH_HIGHLEVEL_SECURITY struct {
	ST_dwSize     DWORD                  // 结构体大小
	ST_szIP       [64]byte               // IP
	ST_nPort      int32                  // 端口
	ST_szUserName [64]byte               // 用户名
	ST_szPassword [64]byte               // 密码
	ST_emSpecCap  EM_LOGIN_SPAC_CAP_TYPE // 登录模式
	ST_byReserved [4]byte                // 字节对齐
	ST_pCapParam  uintptr                // 见CLIENT_LoginEx 接口 pCapParam 与 nSpecCap 关系
}

type NET_DEVICEINFO_Ex struct {
	ST_sSerialNumber    [48]byte // 序列号
	ST_nAlarmInPortNum  int32    // DVR报警输入个数
	ST_nAlarmOutPortNum int32    // DVR报警输出个数
	ST_nDiskNum         int32    // DVR硬盘个数
	ST_nDVRType         int32    // DVR类型，见枚举	NET_DEVICE_TYPE
	ST_nChanNum         int32    // DVR通道个数
	ST_byLimitLoginTime byte     // 在线超时时间，为0表示不限制登录，非0表示限制的分钟数
	ST_byLeftLogTimes   byte     // 当登录失败原因为密码错误时，通过此参数通知用户，剩余登录次数，为0时表示此参数无效
	ST_bReserved        [2]byte  // 保留字节，字节对齐
	ST_nLockLeftTime    int32    // 当登录失败，用户解锁剩余时间（秒数）， -1表示设备未设置该参数
	ST_Reserved         [24]byte // 保留
}

// LoginWithHighLevelSecurity output param
type NET_OUT_LOGIN_WITH_HIGHLEVEL_SECURITY struct {
	ST_dwSize       DWORD             // 结构体大小
	ST_stDeviceInfo NET_DEVICEINFO_Ex // 设备信息
	ST_nError       int32             // 错误码，见CLIENT_Login接口错误码
	ST_byReserved   [132]byte         // 预留字段
}

type SDK_POINT struct {
	ST_nx int16
	ST_ny int16
}

// ListenServer 接口回调fServiceCallBack函数支持的命令类型
const (
	NET_DVR_DISCONNECT          = -1   // Device disconnection callback during the verification period
	NET_DVR_SERIAL_RETURN       = iota // Device send out SN callback char* szDevSerial
	NET_DEV_AUTOREGISTER_RETURN        // when device registering,serial number and token to carry, corresponding NET_CB_AUTOREGISTER
	NET_DEV_NOTIFY_IP_RETURN           // Equipment is only reported IP, not used for active registration
)

// 普通事件报警类型定义
const (
	NET_MOTION_ALARM_EX    = 0x2102 // Motion detection alarm
	NET_EVENT_MOTIONDETECT = 0x218f // Video motion detect event  (Corresponding to structure ALARM_MOTIONDETECT_INFO)
)

// Log print info
type LOG_SET_PRINT_INFO struct {
	ST_dwSize            DWORD                    // Struct size
	ST_bSetFilePath      int32                    // reset log path
	ST_szLogFilePath     [260]byte                // log path(defult"./sdk_log/sdk_log.log")
	ST_bSetFielSize      int32                    // reset log size
	ST_nFileSize         uint32                   // each log file size(default size 10240), unit:KB
	ST_bSetFileNum       int32                    // reset log file number
	ST_nFileNum          uint32                   // log file quanitity(default size 10)
	ST_bSetPrintStrategy int32                    // reset log print strategy
	ST_nPrintStrategy    uint32                   // log out strategy, 0: output to file(default); 1:	output to window
	ST_byReserved        [4]byte                  // Byte alignment
	ST_cbSDKLogCallBack  uintptr                  // log callback, (default NULL)
	ST_dwUser            CallBack_fSDKLogCallBack // UserData
}

type NET_POINT struct {
	ST_nx int16
	ST_ny int16
}

type SDK_RECT struct {
	ST_left   int
	ST_top    int
	ST_right  int
	ST_bottom int
}

type NET_RECT struct {
	ST_nLeft   int32
	ST_nTop    int32
	ST_nRight  int32
	ST_nBottom int32
}

type NET_TIME struct {
	ST_dwYear   DWORD // 年
	ST_dwMonth  DWORD // 月
	ST_dwDay    DWORD // 日
	ST_dwHour   DWORD // 时
	ST_dwMinute DWORD // 分
	ST_dwSecond DWORD // 秒
}

type NET_TIME_EX struct {
	ST_dwYear        DWORD    // 年
	ST_dwMonth       DWORD    // 月
	ST_dwDay         DWORD    // 日
	ST_dwHour        DWORD    // 时
	ST_dwMinute      DWORD    // 分
	ST_dwSecond      DWORD    // 秒
	ST_dwMillisecond DWORD    // 毫秒
	ST_dwUTC         DWORD    // utc时间(获取时0表示无效，非0有效   下发无效)
	ST_dwReserved    [1]DWORD // 预留字段
}

type SDK_EVENT_FILE_INFO struct {
	ST_bCount      BYTE        // 当前文件所在文件组中的文件总数
	ST_bIndex      BYTE        // 当前文件在文件组中的文件编号(编号1开始)
	ST_bFileTag    BYTE        // 文件标签, EM_EVENT_FILETAG
	ST_bFileType   BYTE        // 文件类型,0-普通 1-合成 2-抠图
	ST_stuFileTime NET_TIME_EX // 文件时间
	ST_nGroupId    DWORD       // 同一组抓拍文件的唯一标识
}

// 图片分辨率
type SDK_RESOLUTION_INFO struct {
	ST_snWidth uint16 // 宽
	ST_snHight uint16 // 高
}

type SDK_PIC_INFO struct {
	ST_dwOffSet     DWORD     // 文件在二进制数据块中的偏移位置, 单位:字节
	ST_dwFileLenth  DWORD     // 文件大小, 单位:字节
	ST_wWidth       WORD      // 图片宽度, 单位:像素
	ST_wHeight      WORD      // 图片高度, 单位:像素
	ST_pszFilePath  uintptr   // 鉴于历史原因,该成员只在事件上报时有效// 文件路径// 用户使用该字段时需要自行申请空间进行拷贝保存
	ST_bIsDetected  BYTE      // 图片是否算法检测出来的检测过的提交识别服务器时,// 则不需要再时检测定位抠图,1:检测过的,0:没有检测过
	ST_bReserved    [3]BYTE   // 预留字节数
	ST_nFilePathLen int32     // 文件路径长度 既pszFilePath 用户申请的大小
	ST_stuPoint     SDK_POINT // 小图左上角在大图的位置，使用绝对坐标系
}

type SDK_MSG_OBJECT struct {
	ST_nObjectID               int32         // 物体ID,每个ID表示一个唯一的物体
	ST_szObjectType            [128]byte     // 物体类型
	ST_nConfidence             int32         // 置信度(0~255),值越大表示置信度越高
	ST_nAction                 int32         // 物体动作:1:Appear 2:Move 3:Stay 4:Remove 5:Disappear 6:Split 7:Merge 8:Rename
	ST_BoundingBox             SDK_RECT      // 包围盒
	ST_Center                  SDK_POINT     // 物体型心
	ST_nPolygonNum             int32         // 多边形顶点个数
	ST_Contour                 [16]SDK_POINT // 较精确的轮廓多边形
	ST_rgbaMainColor           DWORD         // 表示车牌、车身等物体主要颜色；按字节表示,分别为红、绿、蓝和透明度,例如:RGB值为(0,255,0),透明度为0时, 其值为0x00ff0000.
	ST_szText                  [128]byte     // 物体上相关的带0结束符文本,比如车牌,集装箱号等等
	ST_szObjectSubType         [62]byte      // 物体子类别,根据不同的物体类型,可以取以下子类型：
	ST_wColorLogoIndex         WORD          // 车标索引
	ST_wSubBrand               WORD          // 车辆子品牌 需要通过映射表得到真正的子品牌 映射表详见开发手册
	ST_byReserved1             BYTE
	ST_bPicEnble               byte         // 是否有物体对应图片文件信息
	ST_stPicInfo               SDK_PIC_INFO // 物体对应图片信息
	ST_bShotFrame              byte         // 是否是抓拍张的识别结果
	ST_bColor                  byte         // 物体颜色(rgbaMainColor)是否可用
	ST_byReserved2             byte
	ST_byTimeType              byte                     // 时间表示类型,详见EM_TIME_TYPE说明
	ST_stuCurrentTime          NET_TIME_EX              // 针对视频浓缩,当前时间戳（物体抓拍或识别时,会将此识别智能帧附在一个视频帧或jpeg图片中,此帧所在原始视频中的出现时间）
	ST_stuStartTime            NET_TIME_EX              // 开始时间戳（物体开始出现时）
	ST_stuEndTime              NET_TIME_EX              // 结束时间戳（物体最后出现时）
	ST_stuOriginalBoundingBox  SDK_RECT                 // 包围盒(绝对坐标)
	ST_stuSignBoundingBox      SDK_RECT                 // 车标坐标包围盒
	ST_dwCurrentSequence       DWORD                    // 当前帧序号（抓下这个物体时的帧）
	ST_dwBeginSequence         DWORD                    // 开始帧序号（物体开始出现时的帧序号）
	ST_dwEndSequence           DWORD                    // 结束帧序号（物体消逝时的帧序号）
	ST_nBeginFileOffset        int64                    // 开始时文件偏移, 单位: 字节（物体开始出现时,视频帧在原始视频文件中相对于文件起始处的偏移）
	ST_nEndFileOffset          int64                    // 结束时文件偏移, 单位: 字节（物体消逝时,视频帧在原始视频文件中相对于文件起始处的偏移）
	ST_byColorSimilar          [NET_COLOR_TYPE_MAX]BYTE // 物体颜色相似度,取值范围：0-100,数组下标值代表某种颜色,详见EM_COLOR_TYPE
	ST_byUpperBodyColorSimilar [NET_COLOR_TYPE_MAX]BYTE // 上半身物体颜色相似度(物体类型为人时有效)
	ST_byLowerBodyColorSimilar [NET_COLOR_TYPE_MAX]BYTE // 下半身物体颜色相似度(物体类型为人时有效)
	ST_nRelativeID             int32                    // 相关物体ID
	ST_szSubText               [20]BYTE                 // "ObjectType"为"Vehicle"或者"Logo"时,表示车标下的某一车系,比如奥迪A6L,由于车系较多,SDK实现时透传此字段,设备如实填写。
	ST_wBrandYear              WORD                     // 车辆品牌年款 需要通过映射表得到真正的年款 映射表详见开发手册
}

type EVENT_CUSTOM_WEIGHT_INFO struct {
	ST_dwRoughWeight DWORD    // 毛重,车辆满载货物重量。单位KG
	ST_dwTareWeight  DWORD    // 皮重,空车重量。单位KG
	ST_dwNetWeight   DWORD    // 净重,载货重量。单位KG
	ST_bReserved     [28]BYTE // 预留字节
}

// 雷达自由流信息
type NET_RADAR_FREE_STREAM struct {
	ST_nABSTime   uint64   // 1年1月1日0时起至今的毫秒数
	ST_nVehicleID int32    // 车辆ID
	ST_unOBUMAC   uint32   // OBU的MAC地址
	ST_byReserved [16]BYTE // 保留字节
}

//卡口事件专用定制上报内容，定制需求增加到Custom下
type EVENT_JUNCTION_CUSTOM_INFO struct {
	ST_stuWeightInfo          EVENT_CUSTOM_WEIGHT_INFO // 原始图片信息
	ST_nCbirFeatureOffset     DWORD                    // 数据偏移，单位字节 （由于结构体保留字节有限的限制,添加在此处， 下同）
	ST_nCbirFeatureLength     DWORD                    // 数据大小，单位字节
	ST_dwVehicleHeadDirection DWORD                    // 车头朝向 0:未知 1:左 2:中 3:右
	ST_byReserved1            [4]BYTE                  // 字节对齐
	ST_stuRadarFreeStream     NET_RADAR_FREE_STREAM    // 雷达自由流信息
	ST_bReserved              [12]BYTE                 // 预留字节
}

// GPS信息
type NET_GPS_INFO struct {
	ST_nLongitude uint32 // 经度(单位是百万分之一度)
	// 西经：0 - 180000000				实际值应为: 180*1000000 – dwLongitude
	// 东经：180000000 - 360000000		实际值应为: dwLongitude – 180*1000000
	// 如: 300168866应为（300168866 - 180*1000000）/1000000 即东经120.168866度
	ST_nLatidude uint32 // 纬度(单位是百万分之一度)
	// 南纬：0 - 90000000				实际值应为: 90*1000000 – dwLatidude
	// 北纬：90000000 – 180000000		实际值应为: dwLatidude – 90*1000000
	// 如: 120186268应为 (120186268 - 90*1000000)/1000000 即北纬30. 186268度
	ST_dbAltitude float64 // 高度,单位为米
	ST_dbSpeed    float64 // 速度,单位km/H
	ST_dbBearing  float64 // 方向角,单位°
	ST_bReserved  [8]BYTE // 保留字段
}

// 颜色RGBA
type NET_COLOR_RGBA struct {
	ST_nRed   int32 // 红
	ST_nGreen int32 // 绿
	ST_nBlue  int32 // 蓝
	ST_nAlpha int32 // 透明
}

// 事件扩展信息
type NET_EXTENSION_INFO struct {
	ST_szEventID  [52]byte // 国标事件ID
	ST_byReserved [80]BYTE // 保留字节
}

type DEV_EVENT_TRAFFICJUNCTION_INFO struct {
	ST_nChannelID         int32                             // 通道号
	ST_szName             [128]byte                         // 事件名称
	ST_byMainSeatBelt     BYTE                              // 主驾驶座,系安全带状态,1-系安全带,2-未系安全带
	ST_bySlaveSeatBelt    BYTE                              // 副驾驶座,系安全带状态,1-系安全带,2-未系安全带
	ST_byVehicleDirection BYTE                              // 当前被抓拍到的车辆是车头还是车尾,具体请见 EM_VEHICLE_DIRECTION
	ST_byOpenStrobeState  BYTE                              // 开闸状态,具体请见 EM_OPEN_STROBE_STATE
	ST_PTS                float64                           // 时间戳(单位是毫秒)
	ST_UTC                NET_TIME_EX                       // 事件发生的时间
	ST_nEventID           int32                             // 事件ID
	ST_stuObject          SDK_MSG_OBJECT                    // 检测到的物体
	ST_nLane              int32                             // 对应车道号
	ST_dwBreakingRule     DWORD                             // 违反规则掩码,第一位:闯红灯;// 第二位:不按规定车道行驶;// 第三位:逆行; 第四位：违章掉头;// 第五位:交通堵塞; 第六位:交通异常空闲// 第七位:压线行驶; 否则默认为:交通路口事件
	ST_RedLightUTC        NET_TIME_EX                       // 红灯开始UTC时间
	ST_stuFileInfo        SDK_EVENT_FILE_INFO               // 事件对应文件信息
	ST_nSequence          int32                             // 表示抓拍序号,如3,2,1,1表示抓拍结束,0表示异常结束
	ST_nSpeed             int32                             // 车辆实际速度Km/h
	ST_bEventAction       BYTE                              // 事件动作,0表示脉冲事件,1表示持续性事件开始,2表示持续性事件结束;
	ST_byDirection        BYTE                              // 路口方向,1-表示正向,2-表示反向
	ST_byLightState       BYTE                              // LightState表示红绿灯状态:0 未知,1 绿灯,2 红灯,3 黄灯
	ST_byReserved         BYTE                              // 保留字节
	ST_byImageIndex       BYTE                              // 图片的序号, 同一时间内(精确到秒)可能有多张图片, 从0开始
	ST_stuVehicle         SDK_MSG_OBJECT                    // 车身信息
	ST_dwSnapFlagMask     DWORD                             // 抓图标志(按位),具体见NET_RESERVED_COMMON
	ST_stuResolution      SDK_RESOLUTION_INFO               // 对应图片的分辨率
	ST_szRecordFile       [128]byte                         // 报警对应的原始录像文件信息
	ST_stuCustomInfo      EVENT_JUNCTION_CUSTOM_INFO        // 自定义信息
	ST_byPlateTextSource  BYTE                              // 车牌识别来源, 0:本地算法识别,1:后端服务器算法识别
	ST_bReserved1         [3]BYTE                           // 保留字节,留待扩展.
	ST_stuGPSInfo         NET_GPS_INFO                      // GPS信息 车载定制
	ST_byNoneMotorInfo    BYTE                              // 0-无非机动车人员信息信息,1-有非机动车人员信息信息// 此字段为1时下面11个字段生效
	ST_byBag              BYTE                              // 是否背包, 0-未知 1-不背包   2-背包
	ST_byUmbrella         BYTE                              // 是否打伞, 0-未知 1-不打伞   2-打伞
	ST_byCarrierBag       BYTE                              // 手提包状态,0-未知 1-没有 2-有
	ST_byHat              BYTE                              // 是否戴帽子, 0-未知 1-不戴帽子 2-戴帽子
	ST_byHelmet           BYTE                              // 头盔状态,0-未知 1-没有 2-有
	ST_bySex              BYTE                              // 性别,0-未知 1-男性 2-女性
	ST_byAge              BYTE                              // 年龄
	ST_stuUpperBodyColor  NET_COLOR_RGBA                    // 上身颜色
	ST_stuLowerBodyColor  NET_COLOR_RGBA                    // 下身颜色
	ST_byUpClothes        BYTE                              // 上身衣服类型 0:未知 1:长袖 2:短袖 3:长裤 4:短裤 5:裙子 6:背心 7:超短裤 8:超短裙
	ST_byDownClothes      BYTE                              // 下身衣服类型 0:未知 1:长袖 2:短袖 3:长裤 4:短裤 5:裙子 6:背心 7:超短裤 8:超短裙
	ST_stuExtensionInfo   NET_EXTENSION_INFO                // 扩展信息
	ST_bReserved          [22]BYTE                          // 保留字节,留待扩展
	ST_nTriggerType       int32                             // TriggerType:触发类型,0车检器,1雷达,2视频,3RSU
	ST_stTrafficCar       DEV_EVENT_TRAFFIC_TRAFFICCAR_INFO // 交通车辆信息
	ST_dwRetCardNumber    DWORD                             // 卡片个数
	ST_stuCardInfo        [16]EVENT_CARD_INFO               // 卡片信息
	ST_stCommInfo         EVENT_COMM_INFO                   // 公共信息
	ST_bNonMotorInfoEx    BOOL                              // 是否有非机动车信息
	ST_stuNonMotor        VA_OBJECT_NONMOTOR                // 非机动车信息
	ST_stuIntelliCommInfo EVENT_INTELLI_COMM_INFO           // 智能事件公共信息
	ST_byReserved2        [1916]BYTE                        // 保留字节,留待扩展
}

// TrafficCar 交通车辆信息
type DEV_EVENT_TRAFFIC_TRAFFICCAR_INFO struct {
	ST_szPlateNumber            [32]BYTE                     // 车牌号码
	ST_szPlateType              [32]BYTE                     // 号牌类型 "Unknown" 未知; "Normal" 蓝牌黑牌; "Yellow" 黄牌; "DoubleYellow" 双层黄尾牌// "Police" 警牌; "Armed" 武警牌; "Military" 部队号牌; "DoubleMilitary" 部队双层// "SAR" 港澳特区号牌; "Trainning" 教练车号牌; "Personal" 个性号牌; "Agri" 农用牌// "Embassy" 使馆号牌; "Moto" 摩托车号牌; "Tractor" 拖拉机号牌; "Other" 其他号牌
	ST_szPlateColor             [32]BYTE                     // 车牌颜色    "Blue","Yellow", "White","Black","YellowbottomBlackText","BluebottomWhiteText","BlackBottomWhiteText","ShadowGreen","YellowGreen"
	ST_szVehicleColor           [32]BYTE                     // 车身颜色    "White", "Black", "Red", "Yellow", "Gray", "Blue","Green"
	ST_nSpeed                   int32                        // 速度    单位Km/H
	ST_szEvent                  [64]BYTE                     // 触发的相关事件    参见事件列表Event List,只包含交通相关事件。
	ST_szViolationCode          [32]BYTE                     // 违章代码    详见TrafficGlobal.ViolationCode
	ST_szViolationDesc          [64]BYTE                     // 违章描述
	ST_nLowerSpeedLimit         int32                        // 速度下限
	ST_nUpperSpeedLimit         int32                        // 速度上限
	ST_nOverSpeedMargin         int32                        // 限高速宽限值    单位：km/h
	ST_nUnderSpeedMargin        int32                        // 限低速宽限值    单位：km/h
	ST_nLane                    int32                        // 车道    参见事件列表Event List中卡口和路口事件。
	ST_nVehicleSize             int32                        // 车辆大小,-1表示未知,否则按位// 第0位:"Light-duty", 小型车// 第1位:"Medium", 中型车// 第2位:"Oversize", 大型车// 第3位:"Minisize", 微型车// 第4位:"Largesize", 长车
	ST_fVehicleLength           float32                      // 车辆长度    单位米
	ST_nSnapshotMode            int32                        // 抓拍方式    0-未分类,1-全景,2-近景,4-同向抓拍,8-反向抓拍,16-号牌图像
	ST_szChannelName            [32]BYTE                     // 本地或远程的通道名称,可以是地点信息    来源于通道标题配置ChannelTitle.Name
	ST_szMachineName            [256]BYTE                    // 本地或远程设备名称    来源于普通配置General.MachineName
	ST_szMachineGroup           [256]BYTE                    // 机器分组或叫设备所属单位    默认为空,用户可以将不同的设备编为一组,便于管理,可重复。
	ST_szRoadwayNo              [64]BYTE                     // 道路编号
	ST_szDrivingDirection       [3][256]BYTE                 //// 行驶方向 , "DrivingDirection" : ["Approach", "上海", "杭州"],// "Approach"-上行,即车辆离设备部署点越来越近；"Leave"-下行,// 即车辆离设备部署点越来越远,第二和第三个参数分别代表上行和// 下行的两个地点
	ST_szDeviceAddress          uintptr                      // 设备地址,OSD叠加到图片上的,来源于配置TrafficSnapshot.DeviceAddress,'\0'结束
	ST_szVehicleSign            [32]BYTE                     // 车辆标识, 例如 "Unknown"-未知, "Audi"-奥迪, "Honda"-本田 ...
	ST_stuSigInfo               SDK_SIG_CARWAY_INFO_EX       // 由车检器产生抓拍信号冗余信息
	ST_szMachineAddr            uintptr                      // 设备部署地点
	ST_fActualShutter           float32                      // 当前图片曝光时间,单位为毫秒
	ST_byActualGain             BYTE                         // 当前图片增益,范围为0~100
	ST_byDirection              BYTE                         // 车道方向,0-南向北 1-西南向东北 2-西向东 3-西北向东南 4-北向南 5-东北向西南 6-东向西 7-东南向西北 8-未知 9-自定义
	ST_byReserved               [2]BYTE                      //
	ST_szDetailedAddress        uintptr                      // 详细地址, 作为szDeviceAddress的补充
	ST_szDefendCode             [64]BYTE                     // 图片防伪码
	ST_nTrafficBlackListID      int32                        // 关联黑名单数据库记录默认主键ID, 0,无效；> 0,黑名单数据记录
	ST_stuRGBA                  SDK_COLOR_RGBA               // 车身颜色RGBA
	ST_stSnapTime               NET_TIME                     // 抓拍时间
	ST_nRecNo                   int32                        // 记录编号
	ST_szCustomParkNo           [33]BYTE                     // 自定义车位号（停车场用）
	ST_byReserved1              [3]BYTE                      //
	ST_nDeckNo                  int32                        // 车板位号
	ST_nFreeDeckCount           int32                        // 空闲车板数量
	ST_nFullDeckCount           int32                        // 占用车板数量
	ST_nTotalDeckCount          int32                        // 总共车板数量
	ST_szViolationName          [64]BYTE                     // 违章名称
	ST_nWeight                  uint32                       // 车重(单位 Kg)
	ST_szCustomRoadwayDirection [32]BYTE                     // 自定义车道方向,byDirection为9时有效
	ST_byPhysicalLane           BYTE                         // 物理车道号,取值0到5
	ST_byReserved2              [3]BYTE                      //
	ST_emMovingDirection        EM_TRAFFICCAR_MOVE_DIRECTION // 车辆行驶方向
	ST_stuEleTagInfoUTC         NET_TIME                     // 对应电子车牌标签信息中的过车时间(ThroughTime)
	ST_stuCarWindowBoundingBox  NET_RECT                     // 车窗包围盒，0~8191
	ST_stuWhiteList             NET_TRAFFICCAR_WHITE_LIST    // 白名单信息
	ST_emCarType                EM_TRAFFICCAR_CAR_TYPE       // 车辆类型
	ST_emLaneType               EM_TRAFFICCAR_LANE_TYPE      // 车道类型
	ST_szVehicleBrandYearText   [64]BYTE                     // 车系年款翻译后文本内容
	ST_szCategory               [32]BYTE                     // 车辆子类型
	ST_stuBlackList             NET_TRAFFICCAR_BLACK_LIST    // 黑名单信息
	ST_bReserved                [240]BYTE                    // 保留字节,留待扩展.
}

// 每个车道的相关信息
type SDK_CARWAY_INFO struct {
	ST_bCarWayID  BYTE                   // 当前车道号
	ST_bReserve   [2]BYTE                // 保留字段
	ST_bSigCount  BYTE                   // 被触发抓拍的个数
	ST_stuSigInfo [3]SDK_SIG_CARWAY_INFO // 当前车道上,被触发抓拍对应的抓拍信息
	ST_bReserved  [12]BYTE               // 保留字段
}

// 白名单权限列表
type NET_WHITE_LIST_AUTHORITY_LIST struct {
	ST_bOpenGate BOOL     // 是否有开闸权限
	ST_bReserved [16]BYTE // 保留字节
}

// 白名单信息
type NET_TRAFFICCAR_WHITE_LIST struct {
	ST_bTrustCar        BOOL                          // 车牌是否属于白名单
	ST_stuBeginTime     NET_TIME                      // 白名单起始时间
	ST_stuCancelTime    NET_TIME                      // 白名单过期时间
	ST_stuAuthorityList NET_WHITE_LIST_AUTHORITY_LIST // 白名单权限列表
	ST_bReserved        [32]BYTE                      // 保留字节
}

// 黑名单信息
type NET_TRAFFICCAR_BLACK_LIST struct {
	ST_bEnable       BOOL     // 是否已启用黑名单
	ST_bIsBlackCar   BOOL     // 车牌是否属于黑名单
	ST_stuBeginTime  NET_TIME // 黑名单起始时间
	ST_stuCancelTime NET_TIME // 黑名单过期时间
	ST_bReserved     [32]BYTE // 保留字节
}

// 非机动车对象
type VA_OBJECT_NONMOTOR struct {
	ST_nObjectID                    int32                            // 物体ID,每个ID表示一个唯一的物体
	ST_emCategory                   EM_CATEGORY_NONMOTOR_TYPE        // 非机动车子类型
	ST_stuBoundingBox               SDK_RECT                         // 包围盒， 非机动车矩形框，0~8191相对坐标
	ST_stuOriginalBoundingBox       SDK_RECT                         // 包围盒， 非机动车矩形框，绝对坐标
	ST_stuMainColor                 NET_COLOR_RGBA                   // 非机动车颜色, RGBA
	ST_emColor                      EM_OBJECT_COLOR_TYPE             // 非机动车颜色, 枚举
	ST_bHasImage                    BOOL                             // 是否有抠图
	ST_stuImage                     NET_NONMOTOR_PIC_INFO            // 物体截图
	ST_nNumOfCycling                int32                            // 骑车人数量
	ST_stuRiderList                 [16]NET_RIDER_INFO               // 骑车人特征,个数和nNumOfCycling关联
	ST_stuSceneImage                SCENE_IMAGE_INFO                 // 全景广角图
	ST_stuFaceSceneImage            FACE_SCENE_IMAGE                 // 人脸全景广角图
	ST_nNumOfFace                   int32                            // 检测到的人脸数量
	ST_fSpeed                       float32                          // 物体速度，单位为km/h
	ST_stuNonMotorFeatureVectorInfo NET_NONMOTOR_FEATURE_VECTOR_INFO // 非机动车特征值数据在二进制数据中的位置信息
	ST_emNonMotorFeatureVersion     EM_FEATURE_VERSION               // 非机动车特征值版本号
	ST_stuNomotorPlateInfo          NET_NONMOTOR_PLATE_INFO          // 非机动车牌信息
	ST_stuObjCenter                 NET_POINT                        // 物体型心(不是包围盒中心), 0-8191相对坐标, 相对于大图
	ST_byReserved                   [3072]BYTE                       // 保留
}

// 非机动车车牌图片信息
type NET_NONMOTOR_PLATE_IMAGE struct {
	ST_nOffset    uint32    // 在二进制数据块中的偏移
	ST_nLength    uint32    // 图片大小
	ST_nWidth     uint32    // 图片宽度
	ST_nHeight    uint32    // 图片高度
	ST_byReserved [512]BYTE // 保留字节
}

// 非机动车配牌信息
type NET_NONMOTOR_PLATE_INFO struct {
	ST_szPlateNumber          [128]BYTE                // 非机动车车牌号
	ST_stuBoundingBox         NET_RECT                 // 包围盒， 非机动车矩形框，0~8191相对坐标
	ST_stuOriginalBoundingBox NET_RECT                 // 包围盒， 非机动车矩形框，绝对坐标
	ST_stuPlateImage          NET_NONMOTOR_PLATE_IMAGE // 非机动车车牌抠图
	ST_emPlateColor           EM_PLATE_COLOR_TYPE      // 车牌颜色
	ST_byReserved             [132]BYTE                // 保留
}

// 非机动车抠图信息
type NET_NONMOTOR_PIC_INFO struct {
	ST_uOffset    uint32    // 在二进制数据块中的偏移
	ST_uLength    uint32    // 图片大小,单位：字节
	ST_uWidth     uint32    // 图片宽度
	ST_uHeight    uint32    // 图片高度
	ST_szFilePath [260]BYTE // 文件路径
	ST_byReserved [512]BYTE // 保留
}

// 骑车人脸图片信息
type RIDER_FACE_IMAGE_INFO struct {
	ST_nOffSet    uint32   // 在二进制数据块中的偏移
	ST_nLength    uint32   // 图片大小,单位字节
	ST_nWidth     uint32   // 图片宽度(像素)
	ST_nHeight    uint32   // 图片高度(像素)
	ST_byReserved [48]BYTE // 保留
}

// 人脸属性
type NET_FACE_ATTRIBUTE_EX struct {
	ST_emSex          EM_SEX_TYPE         // 性别
	ST_nAge           int32               // 年龄,-1表示该字段数据无效
	ST_emComplexion   EM_COMPLEXION_TYPE  // 肤色
	ST_emEye          EM_EYE_STATE_TYPE   // 眼睛状态
	ST_emMouth        EM_MOUTH_STATE_TYPE // 嘴巴状态
	ST_emMask         EM_MASK_STATE_TYPE  // 口罩状态
	ST_emBeard        EM_BEARD_STATE_TYPE // 胡子状态
	ST_nAttractive    int32               // 魅力值, 0未识别，识别时范围1-100,得分高魅力高
	ST_emGlass        EM_HAS_GLASS        // 眼镜
	ST_emEmotion      EM_EMOTION_TYPE     // 表情
	ST_stuBoundingBox SDK_RECT            // 包围盒(8192坐标系)
	ST_emNation       EM_NATION_TYPE      // 民族
	ST_emStrabismus   EM_STRABISMUS_TYPE  // 斜视状态
	ST_byReserved     [64]BYTE            // 保留字节,留待扩展.
}

// 骑车人信息
type NET_RIDER_INFO struct {
	ST_bFeatureValid      BOOL                      // 是否识别到特征信息, TRUE时下面数据才有效
	ST_emSex              EM_SEX_TYPE               // 性别
	ST_nAge               int32                     // 年龄
	ST_emHelmet           EM_NONMOTOR_OBJECT_STATUS // 头盔状态
	ST_emCall             EM_NONMOTOR_OBJECT_STATUS // 是否在打电话
	ST_emBag              EM_NONMOTOR_OBJECT_STATUS // 是否有背包
	ST_emCarrierBag       EM_NONMOTOR_OBJECT_STATUS // 有没有手提包
	ST_emUmbrella         EM_NONMOTOR_OBJECT_STATUS // 是否打伞
	ST_emGlasses          EM_NONMOTOR_OBJECT_STATUS // 是否有带眼镜
	ST_emMask             EM_NONMOTOR_OBJECT_STATUS // 是否带口罩
	ST_emEmotion          EM_EMOTION_TYPE           // 表情
	ST_emUpClothes        EM_CLOTHES_TYPE           // 上衣类型
	ST_emDownClothes      EM_CLOTHES_TYPE           // 下衣类型
	ST_emUpperBodyColor   EM_OBJECT_COLOR_TYPE      // 上衣颜色
	ST_emLowerBodyColor   EM_OBJECT_COLOR_TYPE      // 下衣颜色
	ST_bHasFaceImage      BOOL                      // 是否有骑车人人脸抠图信息
	ST_stuFaceImage       RIDER_FACE_IMAGE_INFO     // 骑车人人脸特写描述
	ST_bHasFaceAttributes BOOL                      // 是否有人脸属性
	ST_stuFaceAttributes  NET_FACE_ATTRIBUTE_EX     // 人脸属性
	ST_emHasHat           EM_HAS_HAT                // 是否戴帽子
	ST_emCap              EM_CAP_TYPE               // 帽类型
	ST_emHairStyle        EM_HAIR_STYLE             // 头发样式
	ST_byReserved         [304]BYTE                 // 保留
}

// intelli event comm info
type EVENT_INTELLI_COMM_INFO struct {
	ST_emClassType EM_CLASS_TYPE // class type
	ST_nPresetID   int32         // Preset ID, value range is 0~255 and when the value is greater than 0 is valied
	ST_bReserved   [124]BYTE     // reserved
}

// 车辆信息，记录了车头、车尾车牌号和车牌颜色
type EVENT_PLATE_INFO struct {
	ST_szFrontPlateNumber [64]BYTE            // 车头车牌号码
	ST_emFrontPlateColor  EM_PLATE_COLOR_TYPE // 车头车牌颜色
	ST_szBackPlateNumber  [64]BYTE            // 车尾车牌号码
	ST_emBackPlateColor   EM_PLATE_COLOR_TYPE // 车尾车牌颜色
	ST_reversed           [128]BYTE           // 保留
}

// 颜色RGBA
type SDK_COLOR_RGBA struct {
	ST_nRed   int32 // 红
	ST_nGreen int32 // 绿
	ST_nBlue  int32 // 蓝
	ST_nAlpha int32 // 透明
}

type EVENT_COMM_INFO struct {
	ST_emNTPStatus                EM_NTP_STATUS                    // NTP校时状态
	ST_nDriversNum                int32                            // 驾驶员信息数
	ST_pstDriversInfo             *SDK_MSG_OBJECT_EX               // 驾驶员信息数据
	ST_pszFilePath                uintptr                          // 本地硬盘或者sd卡成功写入路径,为NULL时,路径不存在
	ST_pszFTPPath                 uintptr                          // 设备成功写到ftp服务器的路径
	ST_pszVideoPath               uintptr                          // 当前接入需要获取当前违章的关联视频的FTP上传路径
	ST_stCommSeat                 [8]EVENT_COMM_SEAT               // 驾驶位信息
	ST_nAttachmentNum             int32                            // 车辆物件个数
	ST_stuAttachment              [8]EVENT_COMM_ATTACHMENT         // 车辆物件信息
	ST_nAnnualInspectionNum       int32                            // 年检标志个数
	ST_stuAnnualInspection        [8]NET_RECT                      // 年检标志
	ST_fHCRatio                   float32                          // HC所占比例，单位：%/1000000
	ST_fNORatio                   float32                          // NO所占比例，单位：%/1000000
	ST_fCOPercent                 float32                          // CO所占百分比，单位：% 取值0~100
	ST_fCO2Percent                float32                          // CO2所占百分比，单位：% 取值0~100
	ST_fLightObscuration          float32                          // 不透光度，单位：% 取值0~100
	ST_nPictureNum                int32                            // 原始图片张数
	ST_stuPicInfos                [6]EVENT_PIC_INFO                // 原始图片信息
	ST_fTemperature               float32                          // 温度值,单位摄氏度
	ST_nHumidity                  int32                            // 相对湿度百分比值
	ST_fPressure                  float32                          // 气压值,单位Kpa
	ST_fWindForce                 float32                          // 风力值,单位m/s
	ST_nWindDirection             uint32                           // 风向,单位度,范围:[0,360]
	ST_fRoadGradient              float32                          // 道路坡度值,单位度
	ST_fAcceleration              float32                          // 加速度值,单位:m/s2
	ST_stuRFIDEleTagInfo          NET_RFIDELETAG_INFO              // RFID 电子车牌标签信息
	ST_stuBinarizedPlateInfo      EVENT_PIC_INFO                   // 二值化车牌抠图
	ST_stuVehicleBodyInfo         EVENT_PIC_INFO                   // 车身特写抠图
	ST_emVehicleTypeInTollStation EM_VEHICLE_TYPE                  // 收费站车型分类
	ST_emSnapCategory             EM_SNAPCATEGORY                  // 抓拍的类型，默认为机动车
	ST_nRegionCode                int32                            // 车牌所属地区代码,(孟加拉海外车牌识别项目),默认-1表示未识别
	ST_emVehicleTypeByFunc        EM_VEHICLE_TYPE_BY_FUNC          // 按功能划分的车辆类型
	ST_emStandardVehicleType      EM_STANDARD_VEHICLE_TYPE         // 标准车辆类型
	ST_nExtraPlateCount           uint32                           // 额外车牌数量
	ST_szExtraPlateNumber         [3][32]BYTE                      // 额外车牌信息
	ST_emOverseaVehicleCategory   EM_OVERSEA_VEHICLE_CATEGORY_TYPE // 海外车辆类型中的子类别
	ST_szProvince                 [64]BYTE                         // 车牌所属国家的省、州等地区名
	ST_bReserved                  [500]BYTE                        // 预留字节
	ST_szCountry                  [20]BYTE                         // 国家
}

// 事件上报携带卡片信息
type EVENT_CARD_INFO struct {
	ST_szCardNumber [36]BYTE // 卡片序号字符串
	ST_bReserved    [32]BYTE // 保留字节,留待扩展.
}

// 违规状态
type EVENT_COMM_STATUS struct {
	ST_bySmoking  BYTE     // 是否抽烟
	ST_byCalling  BYTE     // 是否打电话
	ST_szReserved [14]BYTE // 预留字段
}

// 驾驶位违规信息
type EVENT_COMM_SEAT struct {
	ST_bEnable          BOOL                // 是否检测到座驾信息
	ST_emSeatType       EM_COMMON_SEAT_TYPE // 座驾类型, 0:未识别; 1:主驾驶; 2:副驾驶
	ST_stStatus         EVENT_COMM_STATUS   // 违规状态
	ST_emSafeBeltStatus NET_SAFEBELT_STATE  // 安全带状态
	ST_emSunShadeStatus NET_SUNSHADE_STATE  // 遮阳板状态
	ST_szReserved       [24]BYTE            // 预留字节
}

// 车辆物件
type EVENT_COMM_ATTACHMENT struct {
	ST_emAttachmentType EM_COMM_ATTACHMENT_TYPE // 物件类型
	ST_stuRect          NET_RECT                // 坐标
	ST_bReserved        [20]BYTE                // 预留字节
}

// 交通抓图图片信息
type EVENT_PIC_INFO struct {
	ST_nOffset DWORD // 原始图片偏移，单位字节
	ST_nLength DWORD // 原始图片长度，单位字节
}

// 视频分析物体信息扩展结构体
type SDK_MSG_OBJECT_EX struct {
	ST_dwSize                  DWORD
	ST_nObjectID               int32                       // 物体ID,每个ID表示一个唯一的物体
	ST_szObjectType            [128]BYTE                   // 物体类型
	ST_nConfidence             int32                       // 置信度(0~255),值越大表示置信度越高
	ST_nAction                 int32                       // 物体动作:1:Appear 2:Move 3:Stay 4:Remove 5:Disappear 6:Split 7:Merge 8:Rename
	ST_BoundingBox             SDK_RECT                    // 包围盒
	ST_Center                  SDK_POINT                   // 物体型心
	ST_nPolygonNum             int32                       // 多边形顶点个数
	ST_Contour                 [16]SDK_POINT               // 较精确的轮廓多边形
	ST_rgbaMainColor           DWORD                       // 表示车牌、车身等物体主要颜色；按字节表示,分别为红、绿、蓝和透明度,例如:RGB值为(0,255,0),透明度为0时, 其值为0x00ff0000.
	ST_szText                  [128]BYTE                   // 同SDK_MSG_OBJECT相应字段
	ST_szObjectSubType         [64]BYTE                    // 物体子类别,根据不同的物体类型,可以取以下子类型：// 同SDK_MSG_OBJECT相应字段
	ST_byReserved1             [3]BYTE                     //
	ST_bPicEnble               BYTE                        // 是否有物体对应图片文件信息
	ST_stPicInfo               SDK_PIC_INFO                // 物体对应图片信息
	ST_bShotFrame              BYTE                        // 是否是抓拍张的识别结果
	ST_bColor                  BYTE                        // 物体颜色(rgbaMainColor)是否可用
	ST_bLowerBodyColor         BYTE                        // 下半身颜色(rgbaLowerBodyColor)是否可用
	ST_byTimeType              BYTE                        // 时间表示类型,详见EM_TIME_TYPE说明
	ST_stuCurrentTime          NET_TIME_EX                 // 针对视频浓缩,当前时间戳（物体抓拍或识别时,会将此识别智能帧附在一个视频帧或jpeg图片中,此帧所在原始视频中的出现时间）
	ST_stuStartTime            NET_TIME_EX                 // 开始时间戳（物体开始出现时）
	ST_stuEndTime              NET_TIME_EX                 // 结束时间戳（物体最后出现时）
	ST_stuOriginalBoundingBox  SDK_RECT                    // 包围盒(绝对坐标)
	ST_stuSignBoundingBox      SDK_RECT                    // 车标坐标包围盒
	ST_dwCurrentSequence       DWORD                       // 当前帧序号（抓下这个物体时的帧）
	ST_dwBeginSequence         DWORD                       // 开始帧序号（物体开始出现时的帧序号）
	ST_dwEndSequence           DWORD                       // 结束帧序号（物体消逝时的帧序号）
	ST_nBeginFileOffset        int64                       // 开始时文件偏移, 单位: 字节（物体开始出现时,视频帧在原始视频文件中相对于文件起始处的偏移）
	ST_nEndFileOffset          int64                       // 结束时文件偏移, 单位: 字节（物体消逝时,视频帧在原始视频文件中相对于文件起始处的偏移）
	ST_byColorSimilar          [NET_COLOR_TYPE_MAX]BYTE    // 物体颜色相似度,取值范围：0-100,数组下标值代表某种颜色,详见EM_COLOR_TYPE
	ST_byUpperBodyColorSimilar [NET_COLOR_TYPE_MAX]BYTE    // 上半身物体颜色相似度(物体类型为人时有效)
	ST_byLowerBodyColorSimilar [NET_COLOR_TYPE_MAX]BYTE    // 下半身物体颜色相似度(物体类型为人时有效)
	ST_nRelativeID             int32                       // 相关物体ID
	ST_szSubText               [20]BYTE                    // "ObjectType"为"Vehicle"或者"Logo"时,表示车标下的某一车系,比如奥迪A6L,由于车系较多,SDK实现时透传此字段,设备如实填写。
	ST_nPersonStature          int32                       // 入侵人员身高,单位cm
	ST_emPersonDirection       EM_MSG_OBJ_PERSON_DIRECTION // 人员入侵方向
	ST_rgbaLowerBodyColor      DWORD                       // 使用方法同rgbaMainColor,物体类型为人时有效
}

// RFID 电子车牌标签信息
type NET_RFIDELETAG_INFO struct {
	ST_szCardID              [16]BYTE             // 卡号
	ST_nCardType             int32                // 卡号类型, 0:交通管理机关发行卡, 1:新车出厂预装卡
	ST_emCardPrivince        EM_CARD_PROVINCE     // 卡号省份
	ST_szPlateNumber         [32]BYTE             // 车牌号码
	ST_szProductionDate      [16]BYTE             // 出厂日期
	ST_emCarType             EM_CAR_TYPE          // 车辆类型
	ST_nPower                int32                // 功率,单位：千瓦时，功率值范围0~254；255表示该车功率大于可存储的最大功率值
	ST_nDisplacement         int32                // 排量,单位：百毫升，排量值范围0~254；255表示该车排量大于可存储的最大排量值
	ST_nAntennaID            int32                // 天线ID，取值范围:1~4
	ST_emPlateType           EM_PLATE_TYPE        // 号牌种类
	ST_szInspectionValidity  [16]BYTE             // 检验有效期，年-月
	ST_nInspectionFlag       int32                // 逾期未年检标志, 0:已年检, 1:逾期未年检
	ST_nMandatoryRetirement  int32                // 强制报废期，从检验有效期开始，距离强制报废期的年数
	ST_emCarColor            EM_CAR_COLOR_TYPE    // 车身颜色
	ST_nApprovedCapacity     int32                // 核定载客量，该值<0时：无效；此值表示核定载客，单位为人
	ST_nApprovedTotalQuality int32                // 此值表示总质量，单位为百千克；该值<0时：无效；该值的有效范围为0~0x3FF，0x3FF（1023）表示数据值超过了可存储的最大值
	ST_stuThroughTime        NET_TIME_EX          // 过车时间
	ST_emUseProperty         EM_USE_PROPERTY_TYPE // 使用性质
	ST_szPlateCode           [8]BYTE              // 发牌代号，UTF-8编码
	ST_szPlateSN             [16]BYTE             // 号牌号码序号，UTF-8编码
	ST_szTID                 [64]BYTE             // 标签(唯一标识), UTF-8编码
	ST_bReserved             [40]BYTE             // 保留字节,留待扩展.
}

// 全景广角图
type SCENE_IMAGE_INFO struct {
	ST_nOffSet    uint32   // 在二进制数据块中的偏移
	ST_nLength    uint32   // 图片大小,单位字节
	ST_nWidth     uint32   // 图片宽度(像素)
	ST_nHeight    uint32   // 图片高度(像素)
	ST_byReserved [56]BYTE // 预留字节
}

// 人脸全景图
type FACE_SCENE_IMAGE struct {
	ST_nOffSet    uint32   // 在二进制数据块中的偏移
	ST_nLength    uint32   // 图片大小,单位字节
	ST_nWidth     uint32   // 图片宽度(像素)
	ST_nHeight    uint32   // 图片高度(像素)
	ST_byReserved [56]BYTE // 预留字节
}

// 非机动车特征值数据在二进制数据中的位置信息
type NET_NONMOTOR_FEATURE_VECTOR_INFO struct {
	ST_nOffset    uint32   // 非机动车特征值在二进制数据中的偏移, 单位:字节
	ST_nLength    uint32   // 非机动车特征值数据长度, 单位:字节
	ST_byReserved [32]BYTE // 保留字节
}

// 抓拍信息
type SDK_SIG_CARWAY_INFO struct {
	ST_snSpeed      uint16  // 当前车的速度,km/h
	ST_snCarLength  uint16  // 当前车长,分米为单位
	ST_fRedTime     float32 // 当前车道红灯时间,秒.毫秒
	ST_fCapTime     float32 // 当前车道抓拍时间,秒.毫秒
	ST_bSigSequence BYTE    // 当前抓拍序号
	ST_bType        BYTE    // 当前车道的抓拍类型
	// 0: 雷达高限速;1: 雷达低限速;2: 车检器高限速;3:车检器低限速
	// 4: 逆向;5: 闯红灯;6: 红灯亮;7: 红灯灭;8: 全部抓拍或者卡口
	ST_bDirection  BYTE     // 闯红灯类型:01:左转红灯;02:直行红灯;03:右转红灯
	ST_bLightColor BYTE     // 当前车道的红绿灯状态,0: 绿灯, 1: 红灯, 2: 黄灯
	ST_bSnapFlag   [16]BYTE // 设备产生的抓拍标识
}

// 车检器冗余信息
type SDK_SIG_CARWAY_INFO_EX struct {
	ST_byRedundance [8]BYTE   // 由车检器产生抓拍信号冗余信息
	ST_bReserved    [120]BYTE // 保留字段
}

// Corresponding to StartSearchDevices
type DEVICE_NET_INFO_EX struct {
	ST_iIPVersion          int32     // 4 for IPV4, 6 for IPV6
	ST_szIp                [64]byte  // IP IPV4 like"192.168.0.1" IPV6 like "2008::1/64"
	ST_nPort               int32     // port
	ST_szSubmask           [64]byte  // Subnet mask(IPV6 do not have subnet mask)
	ST_szGateway           [64]byte  // Gateway
	ST_szMac               [40]byte  // MAC address
	ST_szDeviceType        [32]byte  // Device type
	ST_byManuFactory       BYTE      // device manufactory, see EM_IPC_TYPE
	ST_byDefinition        BYTE      // 1-Standard definition 2-High definition
	ST_bDhcpEn             bool      // Dhcp, true-open, false-close
	ST_byReserved1         BYTE      // reserved
	ST_verifyData          [88]byte  // ECC data
	ST_szSerialNo          [48]byte  // serial no
	ST_szDevSoftVersion    [128]byte // soft version
	ST_szDetailType        [32]byte  // device detail type
	ST_szVendor            [128]byte // OEM type
	ST_szDevName           [64]byte  // device name
	ST_szUserName          [16]byte  // user name for log in device(it need be filled when modify device ip)
	ST_szPassWord          [16]byte  // pass word for log in device(it need be filled when modify device ip)
	ST_nHttpPort           uint16    // HTTP server port
	ST_wVideoInputCh       uint16    // count of video input channel
	ST_wRemoteVideoInputCh uint16    // count of remote video input
	ST_wVideoOutputCh      uint16    // count of video output channel
	ST_wAlarmInputCh       uint16    // count of alarm input
	ST_wAlarmOutputCh      uint16    // count of alarm output
	ST_bNewWordLen         int32     // TRUE:szNewPassWord Enable
	ST_szNewPassWord       [64]byte  // pass word for log in device(it need be filled when modify device ip)
	ST_byInitStatus        BYTE      // init status
	// bit0~1：0-old device, can not be init; 1-not init; 2-already init
	// bit2~3：0-old device,reserved; 1-connect to public network disable; 2-connect to public network enable
	// bit4~5：0-old device,reserved; 1-connect to cellphone disable; 2-connect to cellphone enable
	// bit6~7: 0- unknown 1-unsupported reset password 2-support password
	ST_byPwdResetWay BYTE // the way supported for reset password:make sense when the device is init
	// bit0-support reset password by cellphone; bit1-support reset password by mail; bit2-support reset password by XML file;
	// bit3-support reset password by security question; bit4-support reset password by change cellphone
	ST_bySpecialAbility BYTE     // special ability of device ,high eight bit, bit0-2D Code:0 support  1 no support, bit1-PN:0 support  1 no support
	ST_szNewDetailType  [64]byte // device detail type
	ST_bNewUserName     int32    // TRUE:szNewUserName enable
	ST_szNewUserName    [64]byte // new user name for login device(it need be filled when modify device ip)
	ST_byPwdFindVersion BYTE     // password find version, effective when device supports reset password
	// 0-device of old scheme reset password version;1-support reset password by reserved contact;
	// 2-support reset password by change contact;
	ST_szDeviceID        [24]byte // Custom item, do not use for general client;
	ST_dwUnLoginFuncMask DWORD    // function mask before login, Bit0 means wifi config
	ST_szMachineGroup    [64]byte // machine group
	ST_cReserved         [12]byte // reserved
}

// Corresponding to StartSearchDevicesEx
type DEVICE_NET_INFO_EX2 struct {
	ST_stuDevInfo DEVICE_NET_INFO_EX // device net info
	ST_szLocalIP  [64]byte           // local ip
	ST_cReserved  [2048]byte         // reserved
}

// send search type
type EM_SEND_SEARCH_TYPE int32

const (
	EM_SEND_SEARCH_TYPE_MULTICAST_AND_BROADCAST EM_SEND_SEARCH_TYPE = iota // multicast and broadcast search
	EM_SEND_SEARCH_TYPE_MULTICAST                                          // multicast search
	EM_SEND_SEARCH_TYPE_BROADCAST                                          // broadcast search
)

// input param of StartSearchDeviceEx
type NET_IN_STARTSERACH_DEVICE struct {
	ST_dwSize          DWORD                 // struct size
	ST_szLocalIp       [64]byte              // local IP
	ST_cbSearchDevices uintptr               // search device call back
	ST_pUserData       IF_fSearchDevicesCBEx // user data
	ST_emSendType      EM_SEND_SEARCH_TYPE   // send search type
}

// output of StartSearchDeviceEx
type NET_OUT_STARTSERACH_DEVICE struct {
	ST_dwSize DWORD // struct size
}

// Snapshot parameter structure
type SNAP_PARAMS struct {
	ST_Channel   uint32 // Snapshot channel
	ST_Quality   uint32 // Image quality:level 1 to level 6
	ST_ImageSize uint32 // Video size;0:QCIF,1:CIF,2:D1
	ST_mode      uint32 // Snapshot mode;0:request one frame,1:send out requestion regularly,2: Request consecutively
	ST_InterSnap uint32 // Time unit is second.If mode=1, it means send out requestion regularly. The time is valid.
	ST_CmdSerial uint32 // Request serial number
	ST_Reserved  [4]uint32
}

// the input parameters of InitDevAccount
type NET_IN_INIT_DEVICE_ACCOUNT struct {
	ST_dwSize        DWORD     // struct size
	ST_szMac         [40]byte  // mac addr
	ST_szUserName    [128]byte // user name
	ST_szPwd         [128]byte // password
	ST_szCellPhone   [32]byte  // cellphone
	ST_szMail        [64]byte  // mail addr
	ST_byInitStatus  uint8     // this field already abandoned
	ST_byPwdResetWay uint8     // the way supported for reset password:byPwdResetWay value of CLIENT_StartSearchDevices's , CLIENT_SearchDevicesByIPs's callback function and CLIENT_SearchDevices
	// the meaning of this parameter refers to DEVICE_NET_INFO_EX, the value must be same as byPwdResetWay returned by CLIENT_StartSearchDevices,CLIENT_SearchDevicesByIPs and CLIENT_SearchDevices
	// bit0 : 1-support reset password by cellphone, you should set cellphone in szCellPhone if you need to set cellphone
	// bit1 : 1-support reset password by mail, you should set mail address in szMail if you need to set mail address
	ST_byReserved [2]BYTE // reserved
}

// the output parameters of InitDevAccount
type NET_OUT_INIT_DEVICE_ACCOUNT struct {
	ST_dwSize DWORD // struct size:assignment when init the struct
}

// Corresponding to SearchDevicesByIPs
type DEVICE_IP_SEARCH_INFO struct {
	ST_dwSize DWORD         // struct size
	ST_nIpNum int32         // the IPs number for search
	ST_szIP   [256][64]byte // the IPs for search

}

// Type of triggeing motion detection
type EM_MOTION_DETECT_TYPE int

const (
	EM_MOTION_DETECT_UNKNOWN           EM_MOTION_DETECT_TYPE = iota // unknown
	EM_MOTION_DETECT_HUMAN                                          // human
	EM_MOTION_DETECT_VEHICLE                                        // vechicle
	EM_MOTION_DETECT_HUMAN_AND_VEHICLE                              // human and vechicle
)

// Region info of motion detection
type NET_MOTIONDETECT_REGION_INFO struct {
	ST_nRegionID    uint32    // region ID
	ST_szRegionName [64]byte  // region name
	ST_bReserved    [508]byte // reserved
}

// alarm event type SDK_EVENT_MOTIONDETECT (video motion detection event) corresponding data description info
type ALARM_MOTIONDETECT_INFO struct {
	ST_dwSize       uint32
	ST_nChannelID   int32       // channel number
	ST_PTS          float64     // timestamp (unit is millisecond)
	ST_UTC          NET_TIME_EX // event occurrence time
	ST_nEventID     int32       // event ID
	ST_nEventAction int32       // event action, 0 means pulse event, 1 means continuous event begin, 2 means continuous event end;

	ST_nRegionNum         uint32                           // count of region
	ST_stuRegion          [32]NET_MOTIONDETECT_REGION_INFO // region info of motion detection
	ST_bSmartMotionEnable int32                            // smart motion detection is enable or not
	ST_nDetectTypeNum     uint32                           // count of triggeing motion detection type
	ST_emDetectType       [32]EM_MOTION_DETECT_TYPE        // triggeing motion detection type, 当nRegionNum大于0时，one-to-one correspondence with stuRegion if nRegionNum is biger than 0;
	// the type is the first value of emDetectType if nRegionNum is 0
}

type NET_RECORDFILE_INFO struct {
	ST_ch              uint32    // Channel number
	ST_filename        [124]byte // File name
	ST_framenum        uint32    // the total number of file frames
	ST_size            uint32    // File length, unit: Kbyte
	ST_starttime       NET_TIME  // Start time
	ST_endtime         NET_TIME  // End time
	ST_driveno         uint32    // HDD number
	ST_startcluster    uint32    // Initial cluster number
	ST_nRecordFileType uint8     // Recorded file type  0:general record;1:alarm record ;2:motion detection;3:card number record ;4:image ; 19:Pos record ;255:all
	ST_bImportantRecID uint8     // 0:general record 1:Important record
	ST_bHint           uint8     // Document Indexing
	ST_bRecType        uint8     // 0-main stream record 1-sub1 stream record 2-sub2 stream record 3-sub3 stream record
}

// Type of video search
type EM_QUERY_RECORD_TYPE int

const (
	EM_RECORD_TYPE_ALL           EM_QUERY_RECORD_TYPE = iota // All the recorded video
	EM_RECORD_TYPE_ALARM                                     // The video of external alarm
	EM_RECORD_TYPE_MOTION_DETECT                             // The video of dynamic detection alarm
	EM_RECORD_TYPE_ALARM_ALL                                 // All the alarmed video
	EM_RECORD_TYPE_CARD                                      // query by the card number
	EM_RECORD_TYPE_CONDITION                                 // query by condition
	EM_RECORD_TYPE_JOIN                                      // combination query
	EM_RECORD_TYPE_CARD_PICTURE                              // query pictures by the card number, used by HB-U,NVS
	EM_RECORD_TYPE_PICTURE                                   // query pictures, used by HB-U,NVS
	EM_RECORD_TYPE_FIELD                                     // query by field
	EM_RECORD_TYPE_INTELLI_VIDEO                             // Smart record search
	EM_RECORD_TYPE_NET_DATA                                  // query network data, used by Jinqiao Internet Bar
	EM_RECORD_TYPE_TRANS_DATA                                // query the video of serial data
	EM_RECORD_TYPE_IMPORTANT                                 // query the important video
	EM_RECORD_TYPE_TALK_DATA                                 // query the recording file
	EM_RECORD_TYPE_POS                                       // query the pos record
	EM_RECORD_TYPE_INVALID       = 256                       // invalid query type
)

type EM_USEDEV_MODE int

// Audio talk way
const (
	NET_TALK_CLIENT_MODE       EM_USEDEV_MODE = iota // Set client-end mode to begin audio talk
	NET_TALK_SERVER_MODE                             // Set server mode to begin audio talk
	NET_TALK_ENCODE_TYPE                             // Set encode format for audio talk
	NET_ALARM_LISTEN_MODE                            // Set alarm subscribe way
	NET_CONFIG_AUTHORITY_MODE                        // Set user right to realize configuration management
	NET_TALK_TALK_CHANNEL                            // set talking channel(0~MaxChannel-1)
	NET_RECORD_STREAM_TYPE                           // set the stream type of the record for query(0-both main and extra stream,1-only main stream,2-only extra stream)
	NET_TALK_SPEAK_PARAM                             // set speaking parameter,corresponding to NET_SPEAK_PARAM
	NET_RECORD_TYPE                                  // Set by time video playback and download the video file TYPE (see.net RECORD TYPE)
	NET_TALK_MODE3                                   // Set voice intercom parameters of three generations of equipment and the corresponding structure NET TALK the EX
	NET_PLAYBACK_REALTIME_MODE                       // set real time playback function(0-off, 1-on)
	NET_TALK_TRANSFER_MODE                           // Judge the voice intercom if it was a forwarding mode, (corresponding to  NET_TALK_TRANSFER_PARAM)
	NET_TALK_VT_PARAM                                // Set VT Talk param (corresponding to  NET_VT_TALK_PARAM)
	NET_TARGET_DEV_ID                                // set target device identifier for searching system capacity information, (not zero - locate device forwards the information)
	NET_AUDIO_RECORD_LENGTH    = 15                  // set audio record length, corresponding to a int
)

// input parameter of interface GetNMPDevInfo
type NET_IN_GET_NMPDEV_INFO struct {
	ST_dwSize DWORD // struct size
}

// Device information attached to the switch
type NET_NEM_CONNECT_DEVICE_INFO struct {
	ST_szAddress      [40]byte  // Device address or domain name
	ST_szProtocolType [32]byte  // Protocol type of connected device
	ST_szDeviceType   [64]byte  // Equipment type
	ST_szMac          [18]byte  // Device MAC address, colon + uppercase
	ST_szName         [64]byte  // Device Name
	ST_szVendor       [32]byte  // Equipment manufacturer
	ST_byReserved     [126]BYTE // Reserved byte
}

// Neighbor device array
type NET_NEIGHBORS_INFO struct {
	ST_szRemotePortId [64]byte                    // Neighbor device connection port description
	ST_szRemoteDevice [128]byte                   // Unique value of neighbor device
	ST_stuDeviceInfo  NET_NEM_CONNECT_DEVICE_INFO // Device information
	ST_byReserved     [128]BYTE                   // Reserved byte
}

// Device status
type EM_DEV_STATUS int32

const (
	EM_DEV_STATUS_UNKNOWN EM_DEV_STATUS = iota - 1 // Unknown
	EM_DEV_STATUS_OFFLINE                          // Off-line
	EM_DEV_STATUS_ONLINE                           // Online
)

// port connection status
type EM_PORT_STATUS int32

const (
	EM_PORT_STATUS_UNKNOWN EM_PORT_STATUS = iota - 1 // Unknown
	EM_PORT_STATUS_DOWN                              // Down
	EM_PORT_STATUS_UP                                // Up
)

// port transport madia
type EM_MEDIA_TYPE int32

const (
	EM_MEDIA_TYPE_UNKNOWN EM_MEDIA_TYPE = iota // Unknown
	EM_MEDIA_TYPE_COPPER                       // Copper
	EM_MEDIA_TYPE_FIBER                        // Fiber
)

// Port duplex status
type EM_DUPLEX_MODE int32

const (
	EM_DUPLEX_MODE_UNKNOWN    EM_DUPLEX_MODE = iota // Unknown
	EM_DUPLEX_MODE_HALFDUPLEX                       // Halfduplex
	EM_DUPLEX_MODE_FULLDUPLEX                       // FullDuplex
)

// PoE status
type EM_POE_STATUS int32

const (
	EM_POE_STATUS_UNKNOWN     EM_POE_STATUS = iota - 1 // Unknown
	EM_POE_STATUS_NONENABLE                            // Non Enable
	EM_POE_STATUS_ENABLE                               // Enable
	EM_POE_STATUS_UNSUPPORTED                          // Unsupported
)

// Port traffic status
type EM_FLOW_STATUS int32

const (
	EM_FLOW_STATUS_UNKNOWN  EM_FLOW_STATUS = iota - 1 // Unknown
	EM_FLOW_STATUS_NORMAL                             // Normal
	EM_FLOW_STATUS_ABNORMAL                           // Abnoraml
)

// Device port information
type NET_NMPDEVPORT_INFO struct {
	ST_dwLocalPortId      DWORD                   // Device connection port number
	ST_emPoEStatus        EM_POE_STATUS           // PoE status
	ST_szLocalPortDescr   [64]byte                // Device connection port description
	ST_nChassisId         int32                   // Box ID of port
	ST_nSlotId            int32                   // Slot ID of port
	ST_szAliasName        [64]byte                // Device port alias
	ST_dwAggregationId    DWORD                   // Aggregation ID (0 means it does not belong to aggregation）
	ST_dwSpeed            DWORD                   // Port negotiation rate，Unit：mbps
	ST_emStatus           EM_PORT_STATUS          // Port connection status
	ST_emMedia            EM_MEDIA_TYPE           // Port transport media
	ST_emDuplexMode       EM_DUPLEX_MODE          // Port duplex status
	ST_fPoEPower          float32                 // PoE power consmuption
	ST_dwInUcastPkts      uint64                  // Number of unicast received messages
	ST_dwOutUcastPkts     uint64                  // Number of unicast messages send
	ST_dwInBroadcastPkts  uint64                  // Number of broadcast received messages
	ST_dwOutBroadcastPkts uint64                  // Number of broadcast messages
	ST_dwInMulticastPkts  uint64                  // Number of messages received by multicast
	ST_dwOutMulticastPkts uint64                  // Number of multicast messages send
	ST_dwInDiscardPkts    uint64                  // Number of packet loss received
	ST_dwOutDiscardPkts   uint64                  // Number of packet loss send
	ST_nInPausePkts       int64                   // Total number of received pause frames (- 1 indicates illegal value)
	ST_nOutPausePkts      int64                   // Total number of pause frames sent (- 1 indicates illegal value)
	ST_nInFilterPkts      int64                   // Total number of packets filtered (- 1 indicates illegal value)
	ST_dwInOctets         uint64                  // Total bytes received
	ST_dwOutOctets        uint64                  // Total bytes send
	ST_dwInRate           uint64                  // Receive real time rate，Unit：Byte
	ST_dwOutRate          uint64                  // Send real time rate，Unit：Byte
	ST_emFlowStatus       EM_FLOW_STATUS          // Port traffic status
	ST_nRetNeighborsNum   int32                   // Number of neighbor devices array
	ST_stuNeighbors       [255]NET_NEIGHBORS_INFO // Neighbor device arrary
	ST_byReserved         [1024]BYTE              //  Reserved byte
}

// Device information
type NET_NMPDEV_INFO struct {
	ST_szAddress       [40]byte      // Device address or domain name
	ST_szDeviceClass   [64]byte      // Device class
	ST_szDeviceType    [64]byte      // Device type
	ST_szMac           [18]byte      // Device MAC address, colon + uppercase
	ST_byReserved2     [6]BYTE       // byte alignment
	ST_szProtocolType  [32]byte      // Protocol type of connected device
	ST_szVersion       [32]byte      // Device version No
	ST_szSerialNo      [64]byte      // Device serial number
	ST_szName          [64]byte      // Device name
	ST_dwRetPortNum    DWORD         // Number of device return ports
	ST_nPortListMaxNum int32         // Number of port information allocated by users
	ST_pstuPorList     uintptr       // Port information group, corresponding to NET_NMPDEVPORT_INFO
	ST_bPoESupport     int32         // Whether Poe is supported
	ST_fPoETotalPower  float32       // Total power consumption of PoE
	ST_fPoERemainPower float32       // Poe remaining power consumption
	ST_dwUptime        DWORD         // Power on operation time (unit: minute)
	ST_nCPUUsage       int32         // CPU usage (percentage), - 1 is illegal
	ST_emStatus        EM_DEV_STATUS // Device status
	ST_szSubnetMask    [40]byte      // Subnet mask
	ST_szParent        [128]byte     // Unique value of the parent device
	ST_szVendor        [32]byte      // Device manufacturer
	ST_dwPortNumber    DWORD         // Total port number of equipment
	ST_byReserved      [1020]BYTE    // Reserved byte
}

// output parameter of interface  GetNMPDevInfo
type NET_OUT_GET_NMPDEV_INFO struct {
	ST_dwSize        DWORD           // struct size:assignment when init the struct
	ST_byReserved    [4]BYTE         // byte alignment
	ST_stuNMPDevInfo NET_NMPDEV_INFO // Device information
}

// The corresponding parameter when setting log in
type NET_PARAM struct {
	ST_nWaittime            int32   // Waiting time(unit is ms), 0:default 5000ms.
	ST_nConnectTime         int32   // Connection timeout value(Unit is ms), 0:default 1500ms.
	ST_nConnectTryNum       int32   // Connection trial times, 0:default 1.
	ST_nSubConnectSpaceTime int32   // Sub-connection waiting time(Unit is ms), 0:default 10ms.
	ST_nGetDevInfoTime      int32   // Access to device information timeout, 0:default 1000ms.
	ST_nConnectBufSize      int32   // Each connected to receive data buffer size(Bytes), 0:default 250*1024
	ST_nGetConnInfoTime     int32   // Access to sub-connect information timeout(Unit is ms), 0:default 1000ms.
	ST_nSearchRecordTime    int32   // Timeout value of search video (unit ms), default 3000ms
	ST_nsubDisconnetTime    int32   // dislink disconnect time,0:default 60000ms
	ST_byNetType            uint8   // net type, 0-LAN, 1-WAN
	ST_byPlaybackBufSize    uint8   // playback data from the receive buffer size(m),when value = 0,default 4M
	ST_bDetectDisconnTime   uint8   // Pulse detect offline time(second) .When it is 0, the default setup is 60s, and the min time is 2s
	ST_bKeepLifeInterval    uint8   // Pulse send out interval(second). When it is 0, the default setup is 10s, the min internal is 2s.
	ST_nPicBufSize          int32   // actual pictures of the receive buffer size(byte)when value = 0,default 2*1024*1024
	ST_bReserved            [4]BYTE // reserved
}

// APP filter login optimize type, operate |
type EM_OPTTYPE_MOBILE_TYPE int32

const (
	OPTTYPE_MOBILE_DEVICE_ATTR    = 0x00000001 // device info
	OPTTYPE_MOBILE_DEVICE_SN      = 0x00000002 // device sequence
	OPTTYPE_MOBILE_DISK_INFO      = 0x00000004 // hard info
	OPTTYPE_MOBILE_DIGITAL_NUM    = 0x00000008 // digital num
	OPTTYPE_MOBILE_ALARM_IO       = 0x00000010 // Alarm IO number
	OPTTYPE_MOBILE_DEVICE_ABILITY = 0x00000020 // Device ability
	OPTTYPE_MOBILE_STATE_SOFTWARE = 0x00001000 // query software APP option
)

// NetSDK optimize type
type EM_OPTIMIZE_TYPE int32

const (
	EM_OPT_TYPE_DEFAULT         EM_OPTIMIZE_TYPE = iota // default to do nothing
	EM_OPT_TYPE_MOBILE_V1                               // optimized for mobile，conflict with the EM_OPT_TYPE_MOBILE_OPTION option
	EM_OPT_TYPE_P2P_NETPARAM_V1                         // the set for P2P network param,  pParam's type is  NET_PARAM*
	EM_OPT_TYPE_MOBILE_OPTION                           // conflict with the EM_OPT_TYPE_MOBILE_V1 option
	// mobile optimize option, value type int, bit value from low to high, see Enum EM_OPTTYPE_MOBILE_TYPE
)

type DEV_EVENT_TRAFFIC_PARKING_INFO struct {
	ST_nChannelID          int32                             // 通道号
	ST_szName              [128]byte                         // 事件名称
	ST_bReserved1          [4]byte                           // 字节对齐
	ST_PTS                 float64                           // 时间戳(单位是毫秒)
	ST_UTC                 NET_TIME_EX                       // 事件发生的时间
	ST_nEventID            int32                             // 事件ID
	ST_stuObject           SDK_MSG_OBJECT                    // 检测到的物体
	ST_stuVehicle          SDK_MSG_OBJECT                    // 车身信息
	ST_nLane               int32                             // 对应车道号
	ST_stuFileInfo         SDK_EVENT_FILE_INFO               // 事件对应文件信息
	ST_bEventAction        BYTE                              // 事件动作,0表示脉冲事件,1表示持续性事件开始,2表示持续性事件结束;
	ST_reserved            [2]BYTE                           // 保留字节
	ST_byImageIndex        BYTE                              // 图片的序号, 同一时间内(精确到秒)可能有多张图片, 从0开始
	ST_stuStartParkingTime NET_TIME_EX                       // 开始停车时间
	ST_nSequence           int32                             // 表示抓拍序号,如3,2,1,1表示抓拍结束,0表示异常结束(bEventAction=2时此参数有效)
	ST_nAlarmIntervalTime  int32                             // 报警时间间隔,单位:秒。(此事件为连续性事件,在收到第一个此事件之后,若在超过间隔时间后未收到此事件的后续事件,则认为此事件异常结束了)
	ST_nParkingAllowedTime int32                             // 允许停车时长,单位：秒。
	ST_nDetectRegionNum    int32                             // 规则检测区域顶点数
	ST_DetectRegion        [20]SDK_POINT                     // 规则检测区域
	ST_dwSnapFlagMask      DWORD                             // 抓图标志(按位),具体见NET_RESEED_COMMON
	ST_stuResolution       SDK_RESOLUTION_INFO               // 对应图片的分辨率
	ST_bIsExistAlarmRecord BOOL                              // rue:有对应的报警录像; false:无对应的报警录像
	ST_dwAlarmRecordSize   DWORD                             // 录像大小
	ST_szAlarmRecordPath   [256]byte                         // 录像路径
	ST_szFTPPath           [256]byte                         // FTP路径
	ST_stuIntelliCommInfo  EVENT_INTELLI_COMM_INFO           // 智能事件公共信息
	ST_byPreAlarm          BYTE                              // 是否为违章预警图片,0 违章停车事件1 预警事件(预警触发后一定时间，车辆还没有离开，才判定为违章)由于此字段会导致事件含义改变，必须和在平台识别预警事件后，才能有此字段,
	ST_bReserved2          [3]byte                           // 保留字节,留待扩展.
	ST_stuGPSInfo          NET_GPS_INFO                      // GPS信息 车载定制
	ST_bReserved           [228]byte                         // 保留字节,留待扩展.
	ST_stTrafficCar        DEV_EVENT_TRAFFIC_TRAFFICCAR_INFO // 交通车辆信息
	ST_stCommInfo          EVENT_COMM_INFO                   // 公共信息
}
