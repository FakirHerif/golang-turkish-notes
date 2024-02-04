# Concurrency vs Parallelism // eş zamanlılık kavramı ve paralellik kavramı arasındaki farklar 

## Concurrent vs. Parallel:

**Concurrent Execution:** İki görevin başlangıç ve bitiş süreleri birbirine denk geliyorsa, bu görevler "concurrent" olarak adlandırılır. Bu, görevlerin tam olarak aynı anda çalıştığı anlamına gelmez. Bir görevin başladığı ve diğerinin bitmediği bir zaman aralığı olduğunda, bu görevler "concurrent" olur.

**Parallel Execution:** İki görevin başlangıç ve bitiş süreleri sadece örtüştüğünde değil, aynı zamanda bir görevin çalıştığı bir an içinde diğer görevin de çalıştığı durumda "parallel" olarak adlandırılır. Yani, bu durumda gerçek bir aynı anda çalışma söz konusudur.

## Neden Concurrency Kullanmalıyız?

İki görevin sırayla tamamlanması yerine, eş zamanlılık kullanarak bu görevlerin birbirini beklememesi ve işlemci kaynaklarının daha etkili kullanılması mümkündür.

Hesaplamalar arasındaki bekleme sürelerini gizlemek, performansı artırabilir. Örneğin, bir görev bellekten veri beklerken diğer görev çalışabilir, bu da bekleme sürelerini gizler.

## Hardware Mapping

"Hardware mapping", görevlerin hangi donanım üzerinde çalışacağını belirleme sürecidir. Bu, görevlerin hangi çekirdeklerde çalışacağını anlamak anlamına gelir.

Donanım eşleme işlemi genellikle doğrudan programcı tarafından kontrol edilmez. Programcı, görevlerin hangi görevlerin eş zamanlı olarak çalışabileceğini tanımlar, ancak hangi görevlerin gerçekte aynı anda çalışacağı, görevlerin donanıma nasıl eşleştirileceğine bağlıdır.
    
Çoğu programlama dilinde, bu donanım eşleme süreci genellikle işletim sistemleri veya dilin çalışma zamanı tarafından otomatik olarak yönetilir.

## Concurrency'nin Performans Artışı

Tek çekirdekli bir sistemde bile, eş zamanlı programlama kullanmak önemli performans avantajları sağlayabilir. Özellikle, görevlerin belirli olayları beklerken diğer görevlerin çalışması, bekleme sürelerini gizleyerek programın daha hızlı çalışmasına olanak tanır.
    
İşte bu nedenle, paralel donanım olmasa bile, eş zamanlı programlama kullanarak önemli performans artışları elde edilebilir.