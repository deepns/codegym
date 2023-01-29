// Design patterns - Singleton
//
// When to use?
//  1. Need single point of access to a shared resource (e.g. database, cache
//     file etc.)
//  2. Ensure the single point of access by restricting the class to have only
//     one object
//
// A way to use?
//  - Make the default constructor private, so other objects can't instantiate
//    the singleton from within the class
//  - Define a static method that acts as the constructor. Construct the class
//    from the static method.
//
// Notes
//  - Protect the instantiation with a lock to prevent creation of multiple
//    instances when called from multiple threads

#include <memory>
#include <iostream>
#include <map>
#include <thread>

class DBConnection
{
private:
    static DBConnection *instance_;
    static std::mutex mutex_;

    std::map<std::string, std::string> db_;
    std::string name_;

    // Keeping the constructor under private to prevent
    // direct instantiation calls

    DBConnection() {}
    DBConnection(std::string &name) : name_(name) {}

public:
    // Should not be cloneable
    DBConnection(DBConnection &other) = delete;

    // Should not be assignable
    void operator=(const DBConnection &) = delete;

    ~DBConnection() { std::cout << "Closing connection...\n"; }

    static DBConnection *GetInstance(std::string name = "default");

    std::string Get(const std::string &key)
    {
        return db_[key];
    }

    void Put(const std::string &key, const std::string &val)
    {
        db_[key] = val;
    }

    std::string Name() const
    {
        return name_;
    }
};

DBConnection *DBConnection::instance_ = nullptr;
std::mutex DBConnection::mutex_;

DBConnection *DBConnection::GetInstance(std::string name)
{
    std::lock_guard<std::mutex> lock(mutex_);
    if (instance_ == nullptr)
    {
        instance_ = new DBConnection(name);
    }
    return instance_;
}

int main()
{
    std::unique_ptr<DBConnection> db(DBConnection::GetInstance());

    db->Put("foo", "bar");
    db->Put("boo", "baz");
    std::cout << "foo=" << db->Get("foo") << std::endl;

    return 0;
}