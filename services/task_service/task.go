package task_service

type TaskNode struct {
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	Type            string     `json:"type"`
	RunFlag         string     `json:"runFlag"`
	Loc             string     `json:"loc"`
	MaxRetryTimes   int        `json:"maxRetryTimes"`
	RetryInterval   int        `json:"retryInterval"`
	Params          TaskParams `json:"params"`
	PreTasks        string     `json:"preTasks"`
	DepList         []string   `json:"depList"`
	Dependence      string     `json:"dependence"`
	ConditionResult string     `json:"conditionResult"`

	WorkerGroup   string `json:"workerGroup"`
	WorkerGroupId int    `json:"workerGroupId"`
	Timeout       string `json:"timeout"`
}

type TaskParams struct {
	// 资源文件列表
	ResourceList []string `json:"resourceList"`

	// Task 本地保留参数
	LocalParams []string `json:"localParams"`

	RawScript string `json:"rawScript"`
}

type Property struct {
}
