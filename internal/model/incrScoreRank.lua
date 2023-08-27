local score = redis.call("zscore", KEYS[1], ARGV[1])
if not score then
    score = 0
end
score = math.floor(score) + ARGV[2] + ARGV[3]
return redis.call('zadd', KEYS[1], score, ARGV[1])