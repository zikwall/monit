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
