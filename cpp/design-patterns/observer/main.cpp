// Design patterns - Observer
//
// When to use?
// subscribers/observers wanting to be notified register with the publisher.
// publisher notifies the subscriber when there is an event to be notified.
// one-to-many relationship between publisher and subscriber
// similar to pub-sub model.
//
// Notes
// publisher and subscriber loosely coupled
// publisher expects the subscribers to adhere to an interface,
// so publisher can notify the subscribers through appropriate
// method.
//
// A way to use?
// - Subscriber(Observer) and Publisher interfaces implemented with abstract base classes
// - Subscribers registers with the publisher
// - Publisher keeps a list of active subscribers
// - When publisher has something to notify, iterate the subscribers
//   and notify them through the method defined in the subscriber interface

#include <iostream>
#include <list>
#include <memory>
#include <unistd.h>

class ISubscribe {
public:
    virtual void Update(std::string& message) = 0;
    virtual int Id() = 0;
};

class Subscriber final : public ISubscribe {
public:
    Subscriber(int id) : id_(id) { // do something 
    }
    void Update(std::string& message) override {
        std::cout << "subscriber:" << id_
                  <<  " Server sent:" << "message=" << message << std::endl;
    }
    int Id() override {
        return id_;
    };
private:
    int id_;    
};

typedef std::shared_ptr<Subscriber> SubscriberPtr;
typedef std::shared_ptr<ISubscribe> ISubscribePtr;

class IPublish {
public:
    virtual void Subscribe(ISubscribePtr subscriber) = 0;
    virtual void UnSubscribe(ISubscribePtr subscriber) = 0;
    virtual void Notify() = 0;
    
};

class Publisher : public IPublish {
public:
    void Subscribe(ISubscribePtr subscriber) override {
        subscribers.push_back(subscriber);
    }
    void UnSubscribe(ISubscribePtr subscriber) override {
        subscribers.remove(subscriber);
    }
    void Notify() override {
        int subscriber_id = 0;
        // if multithreaded, then subscribers need to be protected
        // when notification is in progress.
        for (auto& subscriber : subscribers) {
            std::cout << "publisher: notifying subscriber=" << subscriber << std::endl;
            std::string message = "hello-" + std::to_string(subscriber->Id());
            subscriber->Update(message);
        }
    }
    void DoSomething() {
        // Notify the clients every 10 seconds 5 times
        for (auto i=0; i < 50; i++) {
            if (i % 10 == 0) {
                sleep(1);
                Notify();
            }
        }
    }
private:
    std::list<ISubscribePtr> subscribers;
};

int main() {
    Publisher p;
    
    ISubscribePtr s1(new Subscriber(101));
    ISubscribePtr s2(new Subscriber(102));
    SubscriberPtr s3 = std::make_shared<Subscriber>(103);

    p.Subscribe(s1);
    p.Subscribe(s2);
    p.Subscribe(s3);

    p.DoSomething();

    p.UnSubscribe(s2);
    p.DoSomething();
    return 0;
}