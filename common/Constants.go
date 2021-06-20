package common

const (
	// 任务保存目录
	JOB_SAVE_DIR = "/pod/jobs/"

	// 任务强杀目录
	JOB_KILLER_DIR = "/pod/killer/"

	// 任务锁目录
	JOB_LOCK_DIR = "/pod/lock/"

	// 服务注册目录
	JOB_WORKER_DIR = "/k8s/workers/"

	// 保存任务事件
	JOB_EVENT_SAVE = 1

	// 删除任务事件
	JOB_EVENT_DELETE = 2

	// 强杀任务事件
	JOB_EVENT_KILL = 3
)