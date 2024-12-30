<script setup>
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

</script>

<template>
    <div class="project-page">
        <section class="dashboard-header pt-5">
            <div class="container mx-auto relative">
                <Navbar />
            </div>
        </section>
        <section class="container mx-auto pt-8">
            <div class="flex justify-between items-center mb-6">
                <div class="w-3/4 mr-6">
                    <h2 class="text-4xl text-gray-900 mb-2 font-medium">Dashboard</h2>
                    <ul class="flex mt-2">
                        <li class="mr-6">
                            <a class="text-gray-800 font-bold" href="#"> Your Projects </a>
                        </li>
                        <li class="mr-6">
                            <nuxt-link
                                class="text-gray-500 hover:text-gray-800"
                                to="/dashboard/transactions"
                            >
                                Your Transactions
                            </nuxt-link>
                        </li>
                    </ul>
                </div>
                <div class="w-1/4 text-right">
                    <nuxt-link
                        to="/dashboard/create"
                        class="bg-orange-button hover:bg-green-button text-white font-bold py-4 px-4 rounded inline-flex items-center"
                    >
                        + Create Campaign
                    </nuxt-link>
                </div>
            </div>
            <hr />
            <div class="block mb-2">
                <div class="w-full lg:max-w-full lg:flex mb-4" v-for="campaign in campaigns" :key="campaign.id">
                    <div
                        class="h-48 lg:h-auto lg:w-48 flex-none bg-cover rounded-t lg:rounded-t-none lg:rounded-l text-center overflow-hidden"
                        style="background-image: url('https://tailwindcss.com/img/card-left.jpg');"
                    ></div>
                    <div
                        class="border-r border-b border-l border-gray-400 lg:border-l-0 lg:border-t lg:border-gray-400 bg-white rounded-b lg:rounded-b-none lg:rounded-r p-8 flex flex-col justify-between leading-normal"
                    >
                        <div class="mb-8">
                            <div class="text-gray-900 font-bold text-xl mb-1">
                                {{ campaign.title }}
                            </div>
                            <p class="text-sm text-gray-600 flex items-center mb-2">
                                Rp. {{ campaign.total }} &middot; {{campaign.accumulated_funds}}
                            </p>
                            <p class="text-gray-700 text-base">
                                {{ campaign.description }}
                            </p>
                        </div>
                        <div class="flex items-center">
                            <NuxtLink
                                :to="`dashboard/projects/${campaign.id}`"
                                class="bg-green-button text-white py-2 px-4 rounded"
                            >
                                Detail
                            </NuxtLink>
                        </div>
                    </div>
                </div>
            </div>
        </section>
        <div class="cta-clip -mt-20"></div>
        <section class="call-to-action bg-purple-progress pt-64 pb-10"></section>
        <Footer />
    </div>    
</template>