// Event tracking functions
function trackPageView() {
    alert("Tracking this page view");
    window.snowplow('trackPageView', 'Example events');
}

function playMix() {
    alert("Playing a mix");
    snowplow('trackStructEvent', {
        category: 'Mixes',
        action: 'Play',
        label: 'MrC/fabric-0503-mix',
        property: '',
        value: 0.0
    });
}

function addProduct() {
    alert("Adding a product to basket");
    snowplow('trackStructEvent', {
        category: 'Checkout',
        action: 'Add',
        label: 'ASO01043',
        property: 'blue:xxl',
        value: 2.0
    });
}

// Whitelist of allowed tab names
const ALLOWED_TABS = ['overview', 'example-events', 'opensearch', 'iglu-server', 'control-plane'];

// Tab navigation function
function showTab(tabName) {
    // Sanitize tabName using whitelist
    if (!ALLOWED_TABS.includes(tabName)) {
        console.warn('Invalid tab name:', tabName);
        tabName = 'overview'; // Default to overview for invalid tabs
    }

    // Hide all tab contents
    const tabContents = document.querySelectorAll('.tab-content');
    tabContents.forEach(content => {
        content.classList.remove('active');
    });

    // Remove active class from all nav tabs
    const navTabs = document.querySelectorAll('.nav-tab');
    navTabs.forEach(tab => {
        tab.classList.remove('active');
    });

    // Show selected tab content
    const selectedTab = document.getElementById(tabName);
    if (selectedTab) {
        selectedTab.classList.add('active');
    }

    // Add active class to corresponding nav tab using data attribute
    navTabs.forEach(tab => {
        if (tab.dataset.tab === tabName) {
            tab.classList.add('active');
        }
    });

    // Update URL hash
    window.location.hash = tabName;
}

// Initialize page on DOM load
window.addEventListener('DOMContentLoaded', function() {
    // Set collector endpoint dynamically based on current address
    const collectorUrl = window.location.host;
    const collectorEndpoint = document.getElementById('collector-endpoint');
    const collectorLink = document.getElementById('collector-link');
    const metricsLink = document.getElementById('metrics-link');

    if (collectorEndpoint) {
        collectorEndpoint.textContent = collectorUrl;
    }
    if (collectorLink) {
        collectorLink.textContent = collectorUrl;
    }
    if (metricsLink) {
        metricsLink.textContent = collectorUrl + '/metrics';
    }

    // Setup event listeners for navigation tabs
    const navTabs = document.querySelectorAll('.nav-tab');
    navTabs.forEach(tab => {
        tab.addEventListener('click', function(e) {
            e.preventDefault();
            const tabName = this.dataset.tab;
            if (tabName) {
                showTab(tabName);
            }
        });
    });

    // Setup event listeners for event buttons
    const trackPageViewBtn = document.getElementById('track-pageview-btn');
    const playMixBtn = document.getElementById('play-mix-btn');
    const addProductBtn = document.getElementById('add-product-btn');

    if (trackPageViewBtn) {
        trackPageViewBtn.addEventListener('click', trackPageView);
    }
    if (playMixBtn) {
        playMixBtn.addEventListener('click', playMix);
    }
    if (addProductBtn) {
        addProductBtn.addEventListener('click', addProduct);
    }

    // Handle initial hash with validation
    const hash = window.location.hash.substring(1);
    if (hash) {
        showTab(hash); // showTab now validates the hash
    }
});
