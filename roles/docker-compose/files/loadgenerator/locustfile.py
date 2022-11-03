from locust import HttpUser, between, task

class GatewayUser(HttpUser):
    wait_time = between(1, 5)

    @task
    def makeRequest(self):
        self.client.get('/ids')
