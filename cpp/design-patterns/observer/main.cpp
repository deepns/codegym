// Observer design pattern
//
// similar to pub-sub model.
// subscribers/observers wanting to be notified register with the publisher.
// publisher notifies the subscriber when there is an event to be notified.
// one-to-many relationship between publisher and subscriber

// publisher and subscriber loosely coupled
// publisher expects the subscribers to adhere to an interface,
// so publisher can notify the subscribers through appropriate
// method.

// Interfaces implemented with aubstract base classes
// TBD
//  Need virtual destructors?
//  typedef pointers or use shared-ptr

#include <iostream>
#include <list>
#include <unistd.h>

class ISubscribe {
public:
    virtual void Update(std::string& message) = 0;
};

class Subscriber : public ISubscribe {
public:
    void Update(std::string& message) override;
};

void Subscriber::Update(std::string& message) {
    std::cout << "Server sent:" << "message=" << message << std::endl;
}

class IPublish {
public:
    virtual void Subscribe(ISubscribe* subscriber) = 0;
    virtual void UnSubscribe(ISubscribe* subscriber) = 0;
    virtual void Notify() = 0;
    
};

class Publisher : public IPublish {
public:
    void Subscribe(ISubscribe* subscriber) override;
    void UnSubscribe(ISubscribe* subscriber) override;
    void Notify() override;
    void DoSomething();
private:
    std::list<ISubscribe *> subscribers;
};

void Publisher::Subscribe(ISubscribe* subscriber) {
    subscribers.push_back(subscriber);
}

void Publisher::UnSubscribe(ISubscribe* subscriber) {
    subscribers.remove(subscriber);
}

void Publisher::Notify() {
    int subscriber_id = 0;
    for (auto& subscriber : subscribers) {
        std::cout << "publisher: notifying subscriber=" << subscriber << std::endl;
        std::string message = "hello-" + std::to_string(++subscriber_id);
        subscriber->Update(message);
    }
}

void Publisher::DoSomething() {
    // Notify the clients every 10
    for (auto i=0; i < 100; i++) {
        if (i % 10 == 0) {
            sleep(1);
            Notify();
        }
    }
}

int main() {
    Publisher p;
    Subscriber s1;
    Subscriber s2;

    p.Subscribe(&s1);
    p.Subscribe(&s2);

    p.DoSomething();

    p.UnSubscribe(&s2);
    p.DoSomething();
    return 0;
}




