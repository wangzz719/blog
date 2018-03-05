# 使用 redis 实现分布式锁
## 获取普通锁
``` python
def acquire_lock(conn, lock_name, acquire_timeout=10):
    identifier = str(uuid.uuid4())
    end = time.time() + acquire_timeout
    while time.time() < end:
        if conn.setnx('lock:' + lock_name, identifier):
            return identifier
        time.sleep(0.001)
    return False
```

## 获取带超时特性的锁
``` python
def release_lock(conn, lock_name, identifier):
    pipe = conn.pipline(True)
    lock_name = 'lock:' + lock_name
    while True:
        try:
            pipe.watch(lock_name)
            if pipe.get(lock_name) == identifier:
                pipe.multi()
                pipe.delete(lock_name)
                pipe.execute()
                return True
            pipe.unwatch()
            break
        except redis.exceptions.WatchError:
            pass
    return False
```
## 释放锁
```python
def acquire_lock_with_timeout(conn, lock_name, acquire_timeout=10, lock_timeout=10):
    identifier = str(uuid.uuid4())
    lock_name = 'lock:' + lock_name
    end = time.time() + acquire_timeout
    lock_timeout = int(math.ceil(lock_timeout))
    while time.time() < end:
        if conn.setnx(lock_name, identifier): # conn.set(lock_name, identifier, ex=lock_timeout, nx=True)
            conn.expire(lock_name, lock_timeout)
            return identifier
        elif not conn.ttl(lock_name):
            conn.expire(lock_name, lock_timeout)
        time.sleep(0.001)
    return False
```