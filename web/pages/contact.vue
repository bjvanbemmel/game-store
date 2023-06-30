<template>
    <div
        class="flex flex-col gap-4"
    >
        <h1
            class="text-xl"
        >
            Contact
        </h1>
        <form class="flex flex-col gap-4 w-96">
            <DefaultTextInput
                class="w-full"
                required
                placeholder="Subject"
            />
            <DefaultTextInput
                class="w-full"
                type="email"
                required
                placeholder="Your e-mail"
            />
            <DefaultTextArea
                class="w-full h-48"
                required
                placeholder="Your message"
            />
            <DefaultButton>
                Submit
            </DefaultButton>
        </form>
    </div>
</template>
