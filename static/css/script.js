window.addEventListener('scroll', function() {
    const scrollPosition = window.scrollY;
    const maxScroll = document.documentElement.scrollHeight - window.innerHeight;
    const scrollPercentage = scrollPosition / maxScroll;

    const horizontalScroll = document.querySelector('.horizontal-scroll');
    const horizontalScrollWidth = horizontalScroll.scrollWidth - window.innerWidth;

    horizontalScroll.style.transform = `translateX(-${scrollPercentage * horizontalScrollWidth}px)`;
});
