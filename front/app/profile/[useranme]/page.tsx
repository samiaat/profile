"use client"

import React, { useState } from "react";
import styles from "./style/profilePage.module.css";

;

//test data UI
const dummyUserData = {
  firstName: "John",
  lastName: "Doe",
  aboutMe: "Passionate developer and designer.",
  avatarUrl: "",
  isPublic: true,
  followers: 120,
  following: 80,
  followState: "Following",
  postCount: 3,
};

const dummyPosts = [
  { id: 1, content: "Hello world!", group: null },
  { id: 2, content: "Another post", group: { name: "React Group" } }
];

const Page = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [activeTab, setActiveTab] = useState('Posts');

  const handleTabClick = (tabId: string) => {
    setActiveTab(tabId);
  };

  const isOwner = true; // juste pour UI
  const basePath = "/images?imageName="; // exemple de base path

  return (
    <div className={`${styles.page} ${isModalOpen ? 'modalOpen' : ''}`}>

      <div className={styles.profileSection}>
        <div className={styles.profileContainer}>
          <div className={styles.ImageContainer}>
            <img
              src="./"
              alt="Profile Picture"
              className={styles.profilePic}
            />
          </div>
        </div>

        <div className={styles.profileInfo}>
          <div className={styles.profileDetails}>
            <h1 className={styles.profileName}>
              {dummyUserData.firstName} {dummyUserData.lastName}
            </h1>
            <p className={styles.profileBio}>{dummyUserData.aboutMe}</p>
          </div>
        </div>
      </div>

      <div className={styles.contentContainer}>
        <div className={styles.navTabs}>
          <div className={`${styles.navTab} ${activeTab === 'About' ? styles.active : ''}`} onClick={() => handleTabClick('About')}>
            About
          </div>
          <div className={`${styles.navTab} ${activeTab === 'Posts' ? styles.active : ''}`} onClick={() => handleTabClick('Posts')}>
            Posts
          </div>
        </div>

        <div className={styles.content}>
          {(!dummyUserData.isPublic && !isOwner && dummyUserData.followState !== "Following") ? (
            <div className={styles.privateLogo}>

            </div>
          ) : (
            <>
              {activeTab === 'About' && <About userData={dummyUserData} />}
              {activeTab === 'Posts' && (
                <div className={styles.postsWrapper}>

                </div>
              )}
            </>
          )}
        </div>
      </div>
    </div>
  );
};

export default Page;

const About = ({ userData }: { userData: any }) => {
  return (
    <div className={styles.about}>
      <p>Email: <span className={styles.data}>john@example.com</span></p>
      <p>Username: <span className={styles.data}>johndoe</span></p>
      <p>Nickname: <span className={styles.data}>Johnny</span></p>
      <p>First Name: <span className={styles.data}>{userData.firstName}</span></p>
      <p>Last Name: <span className={styles.data}>{userData.lastName}</span></p>
      <p>Date of Birth: <span className={styles.data}>01-01-1990</span></p>
      <p>About me: <span className={styles.data}>{userData.aboutMe}</span></p>
    </div>
  );
};
