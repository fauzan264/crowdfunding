<script setup>
    import { ref, onMounted } from "vue";
    import { useRoute } from "vue-router";

    // Dummy Data
    const campaigns = [
        {
            "id": "b9d9357d-453a-4652-a0fc-0a073e070b93",
            "title": "Campaign Title 1",
            "total": "10000",
            "description": "This is the description for campaign 1.",
            "accumulated_funds": "80%"
        },
        {
            "id": "bcc39309-5d8b-445d-8750-e8cadd18d38c",
            "title": "Campaign Title 2",
            "total": "20000",
            "description": "This is the description for campaign 2.",
            "accumulated_funds": "80%"
        },
        {
            "id": "01de8efa-27b2-499c-827b-0d73b7237417",
            "title": "Campaign Title 3",
            "total": "30000",
            "description": "This is the description for campaign 3.",
            "accumulated_funds": "80%"
        },
        {
            "id": "1bc2faad-8ed5-4002-a66a-77b39f9f3a14",
            "title": "Campaign Title 4",
            "total": "40000",
            "description": "This is the description for campaign 4.",
            "accumulated_funds": "80%"
        },
        {
            "id": "7294dbc2-6d42-4f36-9604-084a8feef312",
            "title": "Campaign Title 5",
            "total": "50000",
            "description": "This is the description for campaign 5.",
            "accumulated_funds": "80%"
        },
        {
            "id": "b48a0597-15bd-4373-8e57-77b70218548f",
            "title": "Campaign Title 6",
            "total": "60000",
            "description": "This is the description for campaign 6.",
            "accumulated_funds": "80%"
        },
        {
            "id": "90e4809f-2407-4cfc-ada8-5a244f51e73e",
            "title": "Campaign Title 7",
            "total": "70000",
            "description": "This is the description for campaign 7.",
            "accumulated_funds": "80%"
        },
        {
            "id": "b1058401-5316-4691-b6dd-148d83818d95",
            "title": "Campaign Title 8",
            "total": "80000",
            "description": "This is the description for campaign 8.",
            "accumulated_funds": "80%"
        },
        {
            "id": "97867228-df1b-4952-ac1e-d9208a7aa5c5",
            "title": "Campaign Title 9",
            "total": "90000",
            "description": "This is the description for campaign 9.",
            "accumulated_funds": "80%"
        },
        {
            "id": "443c0b0c-ab70-49f8-a0d5-3cabecb61946",
            "title": "Campaign Title 10",
            "total": "100000",
            "description": "This is the description for campaign 10.",
            "accumulated_funds": "80%"
        }
    ]

    const campaign = ref(null);

    const route = useRoute();
    onMounted(() => {
        const campaignId = route.params.id;
        campaign.value = campaigns.find((c) => c.id === campaignId);
    })
</script>

<template>
    <div class="project-page">
        <section class="dashboard-header pt-5">
            <div class="container mx-auto relative">
                <Header />
            </div>
        </section>
        <section class="container mx-auto pt-8">
            <div class="flex justify-between items-center">
                <div class="w-full mr-6">
                    <h2 class="text-4xl text-gray-900 mb-2 font-medium">Dashboard</h2>
                </div>
            </div>
            <div class="flex justify-between items-center">
                <div class="w-3/4 mr-6">
                    <h3 class="text-2xl text-gray-900 mb-4">Campaign Details</h3>
                </div>
                <div class="w-1/4 text-right">
                    <nuxt-link
                        to="/dashboard/projects/create"
                        class="bg-green-button hover:bg-green-button text-white font-bold px-4 py-1 rounded inline-flex items-center"
                    >
                        Edit
                    </nuxt-link>
                </div>
            </div>
            <div class="block mb-2" v-if="campaign">
                <div class="w-full lg:max-w-full lg:flex mb-4">
                    <div
                        class="border border-gray-400 bg-white rounded p-8 flex flex-col justify-between leading-normal"
                    >
                        <div>
                            <div class="text-gray-900 font-bold text-xl mb-2">
                                {{ campaign.title }}
                            </div>
                            <p class="text-sm font-bold flex items-center mb-1">
                                Description
                            </p>
                            <p class="text-gray-700 text-base">
                                {{ campaign.description }}
                            </p>
                            <p class="text-gray-700 text-base">
                                With N-key rollover (NKRO on wired mode only) you can register
                                as many keys as you can press at once without missing out
                                characters. It allows to use all the same media keys as
                                conventional macOS.
                            </p>
                            <p class="text-sm font-bold flex items-center mb-1 mt-4">
                                What Will Funders Get
                            </p>
                            <ul class="list-disc ml-5">
                                {{ campaign.accumulated_funds }}
                            </ul>
                            <p class="text-sm font-bold flex items-center mb-1 mt-4">
                                Price
                            </p>
                            <p class="text-4xl text-gray-700 text-base">{{ campaign.total }}</p>
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex justify-between items-center">
                <div class="w-3/4 mr-6">
                    <h3 class="text-2xl text-gray-900 mb-4 mt-5">Gallery</h3>
                </div>
                <div class="w-1/4 text-right">
                    <a
                        href="#"
                        class="bg-green-button hover:bg-green-button text-white font-bold px-4 py-1 rounded inline-flex items-center"
                    >
                        Upload
                    </a>
                </div>
            </div>
            <div class="flex -mx-2">
                <div
                    class="relative w-1/4 bg-white m-2 p-2 border border-gray-400 rounded"
                >
                    <figure class="item-thumbnail">
                        <img src="/project-slider-1.jpg" alt="" class="rounded w-full" />
                    </figure>
                </div>
                <div
                    class="relative w-1/4 bg-white m-2 p-2 border border-gray-400 rounded"
                >
                    <figure class="item-thumbnail">
                        <img src="/project-slider-2.jpg" alt="" class="rounded w-full" />
                    </figure>
                </div>
                <div
                    class="relative w-1/4 bg-white m-2 p-2 border border-gray-400 rounded"
                >
                    <figure class="item-thumbnail">
                        <img src="/project-slider-3.jpg" alt="" class="rounded w-full" />
                    </figure>
                </div>
                <div
                    class="relative w-1/4 bg-white m-2 p-2 border border-gray-400 rounded"
                >
                    <figure class="item-thumbnail">
                        <img src="/project-slider-4.jpg" alt="" class="rounded w-full" />
                    </figure>
                </div>
            </div>
            <div class="flex justify-between items-center">
                <div class="w-3/4 mr-6">
                    <h3 class="text-2xl text-gray-900 mb-4 mt-5">
                        Transaction History
                    </h3>
                </div>
            </div>
            <div class="block mb-2">
                <div class="w-full lg:max-w-full lg:flex mb-4">
                    <div
                        class="w-full border border-gray-400 lg:border-gray-400 bg-white rounded p-8 flex flex-col justify-between leading-normal"
                    >
                        <div>
                            <div class="text-gray-900 font-bold text-xl mb-1">
                                Galih Pratama
                            </div>
                            <p class="text-sm text-gray-600 flex items-center mb-2">
                                Rp. 200.000 &middot; 12 September 2020
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </section>
        <div class="cta-clip -mt-20"></div>
        <CallToAction />
        <Footer />
    </div>
</template>