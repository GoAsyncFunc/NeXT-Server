package panel

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 生成模拟用户数据
func generateMockUsers(count int) []UserInfo {
	users := make([]UserInfo, count)
	for i := 0; i < count; i++ {
		users[i] = UserInfo{
			Id:          i + 1,
			Uuid:        fmt.Sprintf("uuid-%d-abcd-efgh-ijkl-%d", i, i*1000),
			SpeedLimit:  100000000, // 100 Mbps
			DeviceLimit: 3,
		}
	}
	return users
}

// 生成模拟在线用户数据
func generateMockOnlineUsers(count int) map[int][]string {
	data := make(map[int][]string, count)
	for i := 0; i < count; i++ {
		// 每个用户1-3个在线IP
		ipCount := (i % 3) + 1
		ips := make([]string, ipCount)
		for j := 0; j < ipCount; j++ {
			ips[j] = fmt.Sprintf("192.168.%d.%d", i/256, i%256)
		}
		data[i+1] = ips
	}
	return data
}

// 生成模拟流量数据
func generateMockTraffic(count int) []UserTraffic {
	traffic := make([]UserTraffic, count)
	for i := 0; i < count; i++ {
		traffic[i] = UserTraffic{
			UID:      i + 1,
			Upload:   int64(i * 1024 * 1024),      // 模拟上传流量
			Download: int64(i * 1024 * 1024 * 10), // 模拟下载流量
		}
	}
	return traffic
}

// 打印内存使用情况
func printMemStats(label string) {
	var m runtime.MemStats
	runtime.GC() // 先GC一次获取更准确的数据
	runtime.ReadMemStats(&m)
	log.Printf("[%s] Alloc=%v MiB, TotalAlloc=%v MiB, Sys=%v MiB, NumGC=%v",
		label,
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC)
}

// 获取详细的内存统计
func getMemStats() (allocMiB, sysMiB uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / 1024 / 1024, m.Sys / 1024 / 1024
}

// 打印完整测试报告
func printTestReport(label string, userCount, connCount int, totalUpload, totalDownload int64, startTime time.Time) {
	elapsed := time.Since(startTime)
	allocMiB, sysMiB := getMemStats()
	numCPU := runtime.NumCPU()
	numGoroutine := runtime.NumGoroutine()

	log.Println("")
	log.Println("==================== 测试报告 ====================")
	log.Printf("测试场景: %s", label)
	log.Printf("用户数量: %d", userCount)
	log.Printf("连接总数: %d", connCount)
	log.Printf("总上传流量: %.2f GB", float64(totalUpload)/1024/1024/1024)
	log.Printf("总下载流量: %.2f GB", float64(totalDownload)/1024/1024/1024)
	log.Printf("内存占用: %d MiB (Alloc) / %d MiB (Sys)", allocMiB, sysMiB)
	log.Printf("CPU核心数: %d", numCPU)
	log.Printf("Goroutine数: %d", numGoroutine)
	log.Printf("耗时: %v", elapsed)
	log.Println("==================================================")
	log.Println("")
}

// 测试100用户场景的内存占用
func TestPerformance_100Users(t *testing.T) {
	printMemStats("开始前")

	// 模拟100个用户
	users := generateMockUsers(100)
	log.Printf("生成 %d 个用户数据", len(users))
	printMemStats("生成用户后")

	// 模拟100个在线用户
	onlineUsers := generateMockOnlineUsers(100)
	log.Printf("生成 %d 个在线用户数据", len(onlineUsers))
	printMemStats("生成在线用户后")

	// 模拟100个用户流量
	traffic := generateMockTraffic(100)
	log.Printf("生成 %d 条流量数据", len(traffic))
	printMemStats("生成流量数据后")

	// 保持引用防止被GC
	_ = users
	_ = onlineUsers
	_ = traffic
}

// 测试1000用户场景的内存占用
func TestPerformance_1000Users(t *testing.T) {
	printMemStats("开始前")

	// 模拟1000个用户
	users := generateMockUsers(1000)
	log.Printf("生成 %d 个用户数据", len(users))
	printMemStats("生成用户后")

	// 模拟1000个在线用户
	onlineUsers := generateMockOnlineUsers(1000)
	log.Printf("生成 %d 个在线用户数据", len(onlineUsers))
	printMemStats("生成在线用户后")

	// 模拟1000个用户流量
	traffic := generateMockTraffic(1000)
	log.Printf("生成 %d 条流量数据", len(traffic))
	printMemStats("生成流量数据后")

	_ = users
	_ = onlineUsers
	_ = traffic
}

// ==================== 基准测试 ====================

// 基准测试：在线用户处理 (100用户)
func BenchmarkOnlineUsers_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateMockOnlineUsers(100)
	}
}

// 基准测试：在线用户处理 (1000用户)
func BenchmarkOnlineUsers_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateMockOnlineUsers(1000)
	}
}

// 基准测试：流量数据处理 (100用户)
func BenchmarkTraffic_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateMockTraffic(100)
	}
}

// 基准测试：流量数据处理 (1000用户)
func BenchmarkTraffic_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateMockTraffic(1000)
	}
}

// 基准测试：用户数据处理 (100用户)
func BenchmarkUsers_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateMockUsers(100)
	}
}

// 基准测试：用户数据处理 (1000用户)
func BenchmarkUsers_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateMockUsers(1000)
	}
}

// ==================== 真实场景模拟测试 ====================

// 注意：sync, sync/atomic, time 已在顶部导入

// 模拟单个连接的状态
type MockConnection struct {
	ID         uint64
	UserID     int
	RemoteAddr string
	LocalAddr  string
	StartTime  time.Time
	Upload     int64
	Download   int64
	Active     bool
}

// 模拟用户连接管理器
type MockConnectionManager struct {
	connections sync.Map // map[uint64]*MockConnection
	connCount   atomic.Int64
	userConns   sync.Map // map[int][]uint64 - 用户ID -> 连接ID列表
}

// 创建新连接
func (m *MockConnectionManager) AddConnection(userID int, connID uint64) *MockConnection {
	conn := &MockConnection{
		ID:         connID,
		UserID:     userID,
		RemoteAddr: fmt.Sprintf("192.168.%d.%d:%d", userID/256, userID%256, 10000+int(connID%50000)),
		LocalAddr:  fmt.Sprintf("0.0.0.0:%d", 443),
		StartTime:  time.Now(),
		Upload:     0,
		Download:   0,
		Active:     true,
	}
	m.connections.Store(connID, conn)
	m.connCount.Add(1)

	// 添加到用户连接列表
	if val, ok := m.userConns.Load(userID); ok {
		conns := val.([]uint64)
		m.userConns.Store(userID, append(conns, connID))
	} else {
		m.userConns.Store(userID, []uint64{connID})
	}
	return conn
}

// 模拟流量
func (m *MockConnectionManager) SimulateTraffic(connID uint64, upload, download int64) {
	if val, ok := m.connections.Load(connID); ok {
		conn := val.(*MockConnection)
		atomic.AddInt64(&conn.Upload, upload)
		atomic.AddInt64(&conn.Download, download)
	}
}

// 获取用户总流量
func (m *MockConnectionManager) GetUserTraffic(userID int) (upload, download int64) {
	if val, ok := m.userConns.Load(userID); ok {
		conns := val.([]uint64)
		for _, connID := range conns {
			if connVal, ok := m.connections.Load(connID); ok {
				conn := connVal.(*MockConnection)
				upload += atomic.LoadInt64(&conn.Upload)
				download += atomic.LoadInt64(&conn.Download)
			}
		}
	}
	return
}

// 获取统计信息
func (m *MockConnectionManager) Stats() (totalConns int64, activeUsers int) {
	totalConns = m.connCount.Load()
	m.userConns.Range(func(key, value interface{}) bool {
		activeUsers++
		return true
	})
	return
}

// 测试：100用户，每用户100连接 (共10000连接)
func TestRealistic_100Users_100Conns(t *testing.T) {
	startTime := time.Now()
	printMemStats("开始前")

	mgr := &MockConnectionManager{}
	var connID uint64 = 0

	// 创建100个用户，每个100个连接
	for userID := 1; userID <= 100; userID++ {
		for c := 0; c < 100; c++ {
			connID++
			mgr.AddConnection(userID, connID)
		}
	}

	totalConns, activeUsers := mgr.Stats()
	log.Printf("创建完成: %d 个用户, %d 个连接", activeUsers, totalConns)
	printMemStats("创建连接后")

	// 模拟流量 (每个连接100KB上传, 1MB下载)
	var totalUpload, totalDownload int64
	for i := uint64(1); i <= connID; i++ {
		mgr.SimulateTraffic(i, 1024*100, 1024*1000)
		totalUpload += 1024 * 100
		totalDownload += 1024 * 1000
	}
	printMemStats("模拟流量后")

	// 获取用户流量
	for userID := 1; userID <= 100; userID++ {
		up, down := mgr.GetUserTraffic(userID)
		if userID <= 3 {
			log.Printf("用户 %d: 上传=%d KB, 下载=%d KB", userID, up/1024, down/1024)
		}
	}

	// 打印测试报告
	printTestReport("100用户 x 100连接", activeUsers, int(totalConns), totalUpload, totalDownload, startTime)
}

// 测试：100用户，每用户1000连接 (共100000连接)
func TestRealistic_100Users_1000Conns(t *testing.T) {
	startTime := time.Now()
	printMemStats("开始前")

	mgr := &MockConnectionManager{}
	var connID uint64 = 0

	// 创建100个用户，每个1000个连接
	for userID := 1; userID <= 100; userID++ {
		for c := 0; c < 1000; c++ {
			connID++
			mgr.AddConnection(userID, connID)
		}
	}

	totalConns, activeUsers := mgr.Stats()
	log.Printf("创建完成: %d 个用户, %d 个连接", activeUsers, totalConns)
	printMemStats("创建10万连接后")

	// 模拟流量
	var totalUpload, totalDownload int64
	for i := uint64(1); i <= connID; i++ {
		mgr.SimulateTraffic(i, 1024*100, 1024*1000)
		totalUpload += 1024 * 100
		totalDownload += 1024 * 1000
	}
	printMemStats("模拟流量后")

	// 打印测试报告
	printTestReport("100用户 x 1000连接", activeUsers, int(totalConns), totalUpload, totalDownload, startTime)
}

// 测试：1000用户，每用户1000连接 (共1000000连接 - 百万级)
func TestRealistic_1000Users_1000Conns(t *testing.T) {
	startTime := time.Now()
	printMemStats("开始前")

	mgr := &MockConnectionManager{}
	var connID uint64 = 0

	// 创建1000个用户，每个1000个连接
	for userID := 1; userID <= 1000; userID++ {
		for c := 0; c < 1000; c++ {
			connID++
			mgr.AddConnection(userID, connID)
		}
	}

	totalConns, activeUsers := mgr.Stats()
	log.Printf("创建完成: %d 个用户, %d 个连接", activeUsers, totalConns)
	printMemStats("创建百万连接后")

	// 模拟流量
	var totalUpload, totalDownload int64
	for i := uint64(1); i <= connID; i++ {
		mgr.SimulateTraffic(i, 1024*100, 1024*1000)
		totalUpload += 1024 * 100
		totalDownload += 1024 * 1000
	}
	printMemStats("模拟流量后")

	// 打印测试报告
	printTestReport("1000用户 x 1000连接", activeUsers, int(totalConns), totalUpload, totalDownload, startTime)
}

// 基准测试：创建连接性能
func BenchmarkCreateConnection(b *testing.B) {
	mgr := &MockConnectionManager{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mgr.AddConnection(i%100+1, uint64(i))
	}
}

// 基准测试：流量统计性能
func BenchmarkTrafficStats(b *testing.B) {
	mgr := &MockConnectionManager{}
	// 预创建1000个连接
	for i := 0; i < 1000; i++ {
		mgr.AddConnection(i%100+1, uint64(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mgr.SimulateTraffic(uint64(i%1000), 1024, 1024)
	}
}

// 基准测试：获取用户流量性能
func BenchmarkGetUserTraffic(b *testing.B) {
	mgr := &MockConnectionManager{}
	// 预创建100用户，每用户100连接
	for userID := 1; userID <= 100; userID++ {
		for c := 0; c < 100; c++ {
			mgr.AddConnection(userID, uint64(userID*1000+c))
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mgr.GetUserTraffic(i%100 + 1)
	}
}
