local pipeline = import 'pipeline.libsonnet';
local name = 'drone-apex-up';

[
  pipeline.test,
  pipeline.build(name, 'linux', 'amd64'),
  // pipeline.build(name, 'linux', 'arm64'),
  // pipeline.build(name, 'linux', 'arm'),
  pipeline.release,
  pipeline.notifications(depends_on=[
    'linux-amd64',
    // 'linux-arm64',
    // 'linux-arm',
    'release-binary',
  ]),
  pipeline.secret('docker_username', 'drone/data/docker', 'username'),
  pipeline.secret('docker_password', 'drone/data/docker', 'password'),
  pipeline.secret('api_key', 'drone/data/github', 'api_key'),
]
