import React, { useState } from "react";
import Footer from "../../../Layout/Footer";
import Navbar from "../../../Layout/Navbar";
import Sidebar from "../../../Layout/Sidebar";
import HeroSection from "./../HeroSection";
import InfoSectionDetail from "./InfoSectionDetail";
import { homeObjOne, homeObjTwo, homeObjThree } from "./../InfoSection/Data";

const KaryaTulisDetail = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      <Sidebar isOpen={isOpen} toggle={toggle} />
      <Navbar toggle={toggle} />
      <HeroSection />
      <InfoSectionDetail {...homeObjTwo} />
      <Footer />
    </>
  );
};

export default KaryaTulisDetail;
