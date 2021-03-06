[![build](https://github.com/zikwall/monit/workflows/golangci-lint/badge.svg)](https://github.com/zikwall/monit/actions)

<div align="center">
  <img width="400" height="200" src="https://github.com/zikwall/monit/blob/master/images/1.png">
</div>

### Communication draw

```shell
                                                      (include metrics, http, screenshots)
  | <------- (signals)------> Monitoring <-- (read settings from repository, write data to storage) --> |
  |                                                                                                     |
 API <---- (read/wrire) ----> Repository <----> MySQL (settings, users) <-----------------------------> |
  |      \-- (read/write) --> Storage <-------> Clickhouse (metrcis, heatmap, events, statuses) <-----> |
  |    
 Dashboard (chars, settings, auth)
```

### Services list

- [ ] **API Server**
  - [ ] POST /api/data/heatmap/v1
  - [ ] POST /api/data/events/v1
  - [ ] POST /api/state/streams/v1/[:id]/up
  - [ ] POST /api/state/streams/v1/[:id]/down
  - [ ] DELETE /api/state/streams/v1/[:id]/delete
  - [ ] GET  /api/state/streams/v1
  - [ ] GET  /api/state/streams/v1/[:id]
  - [ ] POST /api/state/streams/v1/[:id]/update
  - [ ] POST /api/state/streams/v1/[:id]/create
  - [ ] POST /api/config/user/v1/[:id]/update
  - [ ] POST /api/config/user/v1/[:id]/create
  - [ ] GET  /api/config/user/v1/[:id]
  - [ ] POST /api/auth/jwt/v1
  - [ ] GET  /api/query/unique/v1
  - [ ] GET  /api/query/platform/v1
  - [ ] GET  /api/query/summary/v1
  - [ ] GET  ... more query like api endpoints
  - [ ] WebSockets for events and heatmaps
  - [ ] Protect all endpoints (with JWT or basic auth or just simple tokens)
  - [ ] Health check API
  - [ ] Metrics API
- [ ] **Storage Server**
  - [ ] WriteHeatmap()
  - [ ] WriteEvent()
  - [ ] Queries
  - [ ] Health check API
  - [ ] Metrics API
- [ ] **Repository Server**
  - [ ] FindStream()
  - [ ] UpdateStream()
  - [ ] DeleteStream()
  - [ ] UpStream()
  - [ ] DownStream()
  - [ ] more like as api enpoints
  - [ ] Health check API
  - [ ] Metrics API
- [ ] **Monitoring Server(-s)**
  - [ ] Perform(stream)
  - [ ] Revoke(stream)
  - [ ] Monitor stream data (with transcode)
  - [ ] Monitor stream screens (with transcode)
  - [ ] Event Loop
  - [ ] Periodical Loop
  - [ ] Health check API
  - [ ] Metrics API

### Describe Tables 

#### (metrics)

| Columns               |     Type      |
|-----------------------|:-------------:|
| **stream_id**         | String        | 
| **bitrate**           | Float64       |   
| **frames**            | UInt64        |   
| **height**            | UInt64        |   
| **fps**               | Float64       |    
| **bytes**             | UInt64        |    
| **seconds**           | Float64       |    
| **keyframe_interval** | UInt64        |    
| **insert_date**       | Date          |    
| **insert_ts**         | DateTime      |

#### (heatmap)

| Columns               |     Type      |
|-----------------------|:-------------:|
| **stream_id**         | String        |
| **uniq_id**           | String        |
| **platform**          | String        |   
| **app**               | String        |   
| **version**           | String        |   
| **os**                | String        |    
| **browser**           | String        |    
| **region**            | String        |    
| **host**              | String        |    
| **user_agent**        | String        |  
| **insert_date**       | Date          |    
| **insert_ts**         | DateTime      |
