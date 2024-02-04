# Concurrency Basics // Eş Zamanlılık Temelleri

## İşlemler (Processes)

Bir işlem, çalışan bir program örneğidir. Her işlemin kendi özel belleği vardır, yani sanal adres alanı, kod, yığın (stack), heap, paylaşılan kütüphaneler ve işleme özgü register'ları bulunur. İşletim sistemi, işlemlerin eş zamanlı olarak çalışmasını sağlar. İşletim sistemi, işlemlere adil bir şekilde işlemci (CPU) kaynaklarına erişim sağlamak ve çakışmaları önlemekle görevlidir.

## Zamanlama (Scheduling)

İşletim sistemi, işlemleri adil bir şekilde işlemciye zaman dilimleri atayarak zamanlar. Zamanlama algoritmaları farklı işletim sistemlerinde değişebilir. Örneğin, yuvarlak-robin algoritması, her işleme eşit zaman dilimleri sağlar. Öncelikli işlemler veya belirli görevler için özel algoritmalar da kullanılabilir.

## İş Parçacıkları ve Gorutinler (Threads and Goroutines)

İş parçacığı (thread), bir işlemin daha hafif bir versiyonudur ve diğer iş parçacıkları ile bazı bağlamları paylaşabilir. İşletim sistemleri genellikle iş parçacıklarını yönetir ve bir işlemden diğerine geçiş yapmak "bağlam değişimi" olarak adlandırılır. Go dilinde ise "gorutinler" kullanılır. Gorutinler, bir işletim sistemi iş parçacığında çalışabilir ve Go Runtime Scheduler tarafından yönetilir. Bu, paralel değil, eş zamanlı bir davranış sağlar.