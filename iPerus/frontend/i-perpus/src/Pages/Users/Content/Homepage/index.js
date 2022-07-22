import React, { useState } from "react";
import Footer from "../../Layout/Footer";
import HeroSection from "./components/HeroSection";
import InfoSection from "./components/InfoSection";
import {
  homeObjOne,
  homeObjTwo,
  homeObjThree,
} from "./components/InfoSection/Data";
import Navbar from "../../Layout/Navbar";
import Services from "./components/Services";
import Sidebar from "../../Layout/Sidebar";

const Home = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      <Sidebar isOpen={isOpen} toggle={toggle} />
      <Navbar toggle={toggle} />
      <HeroSection />
      <InfoSection {...homeObjOne} />
      <Services />
      <Footer />
    </>
  );
};

export default Home;
