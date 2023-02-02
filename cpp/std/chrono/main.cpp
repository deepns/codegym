#include <chrono>
#include <iostream>

int main() {
    std::chrono::seconds min_timeout(5);
    std::chrono::seconds max_timeout(100);

    std::cout << (max_timeout - min_timeout).count() << "\n";

    std::chrono::duration<int> timeout_secs(10);
    std::chrono::duration<int, std::milli> timeout_msecs(10000);
    std::cout << "timeout_secs=" << timeout_secs.count() << "\n";
    std::cout << "timeout_msecs=" << timeout_msecs.count() << "\n";

    timeout_secs = min_timeout;
    std::cout << "timeout_secs=" << timeout_secs.count() << "\n";

    return 0;
}